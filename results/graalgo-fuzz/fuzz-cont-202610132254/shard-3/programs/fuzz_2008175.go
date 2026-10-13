package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int32) int8 {
    v6 := []int{128, 773, 4294967295}
    if false {
        if len(v6) > 0 {
            v6[0] = -1
        }
        v6 = append(v6, 42)
    } else {
        var v7 int32 = p0
        _ = len(v6)
        _ = v7
    }
    fmt.Println(p0)
    if (255 != 0) {
        if true {
            _ = len(v6)
            fmt.Println(p0)
        }
        var v8 int = 32767
        _ = v8
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v9 uint32 = 2147483647
        _ = v9
    }
    _ = p0
    _ = v6
    return 127
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 bool = (9223372036854775807 > 32767)
    var v4 int32 = 0
    var v5 uint32 = 2147483648
    switch "" {
    case "hello":
        v10 := GenS1{F0: -623}
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                v5 += 1
                _ = fi12
            }
            v4 = v4
            _ = fi11
        }
    case "a":
        v13 := GenS1{F0: -533}
        _ = v13
    default:
        fmt.Println(v3, v3)
    }
    delete(v2, 4294967295)
    var v14 int32 = 0
    switch ("foo bar" + "a") {
    case "":
        var v15 uint64 = 70
        _ = v15
    case "hello":
        if v3 {
            if v3 {
                if v3 {
                    v5 = (v5 ^ v5)
                    v3 = (32767 != -891)
                    if (!v3) {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            switch -2147483648 {
                            case 128:
                                var v17 uint = 42
                                _ = v17
                            case -184:
                                v14 += (v14 & 2147483647)
                            case 4294967295:
                                v3 = false
                            }
                            v3 = ((^(fi16)) > fi16)
                            _ = fi16
                        }
                        v18 := []int{2147483647, -2147483648, -2147483648}
                        var v19 uint16 = uint16(v5)
                        _ = v18
                        _ = v19
                    }
                } else {
                    v5 += (^(v5))
                }
            }
        } else {
            v4 -= (-(v4))
            v20 := GenS1{F0: 2147483647}
            var v21 int32 = (v14 | (v4 + (v4 & -742)))
            _ = v20
            _ = v21
        }
    default:
        switch "foo bar" {
        case "\n\t":
            var v22 bool = false
            _ = v22
        }
    }
    var v23 uint64 = 956
    fmt.Println(v4, v14, v14)
    _ = len(v2)
    if true {
        if true {
            v5 = (v5 >> ((416) & 63))
        } else {
            if v3 {
                var v24 uint16 = 111
                fmt.Println(v23, v4)
                switch "a" {
                case "a":
                    fmt.Println(v24, v24)
                case "foo bar":
                    switch 42 {
                    case 592:
                        v25 := GenS1{F0: 42}
                        _ = v25
                    case 0:
                        var v26 uint32 = 0
                        _ = v26
                    default:
                        var v27 uint = 65535
                        _ = v27
                    }
                case "":
                    v23 -= 639
                }
                _ = v24
            } else {
                for fi28 := 0; fi28 < 6; fi28++ {
                    if (false && (!((100 >= 1) && false))) {
                        v5 += (^(v5))
                    }
                    _ = fi28
                }
                v29 := []int{-1}
                if ((0 != 2147483647) && (false || (!true))) {
                    if v3 {
                        var v30 int8 = 42
                        _ = v30
                    } else {
                        v3 = false
                    }
                }
                _ = v29
            }
            var v31 bool = (((1e10 / 1e10) / 1e10) < 1.5)
            v32 := []int{138, 1, 2147483647}
            _ = v31
            _ = v32
        }
    }
    _ = len(v2)
    v33 := []int{2147483647, 128}
    if true {
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 bool = true
            delete(v2, (v5 ^ v5))
            _ = fi34
            _ = v35
        }
    }
    switch 32767 {
    case -92:
        fmt.Println(v14, v14, v5)
    case 4294967295:
        fmt.Println(v14, v23)
    case 9223372036854775807:
        var v36 int32 = -2147483648
        _ = v36
    }
    fmt.Println(v23, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
