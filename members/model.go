package members

type FindMemberServiceOptions struct {
	IsEnabled *bool
	IsDeleted *bool
}

type GetMemberPageServiceOptions struct {
	TagIDs []*string
}

type GetMemberAfterServiceOptions struct {
	Sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetMemberBeforeServiceOptions struct {
	Sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetAllMemberServiceOptions struct {
}
