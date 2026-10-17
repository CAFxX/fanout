package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int8; F2 bool }
func genFn0(p0 uint16) int8 {
    p0 = 676
    if true {
        var v5 int8 = -128
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 float32 = 0.1
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        var v8 bool = true
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int = 1
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(p0, p0)
        switch 2147483647 {
        case -1:
            var v12 uint16 = (p0 & p0)
            _ = v12
        case -9223372036854775808:
            v13 := GenS1{F0: 42, F1: false, F2: (-1 < 42)}
            _ = v13
        case 628:
            var v14 uint64 = 42
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 string = "hello"
    v4 = (v4 + v4)
    if false {
        switch -1 {
        case 42:
            v4 = v4
        case 9223372036854775807:
            switch v4 {
            case "":
                if true {
                    v4 = v4
                } else {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v4 = v4
                        _ = fi15
                    }
                }
            }
        case -1:
            v16 := v1[1]
            _ = v16
        }
        v17 := GenS1{F0: 1, F1: true, F2: false}
        switch ((v4 + "foo bar") + v4) {
        case "":
            v18 := GenS2{F0: 894}
            _ = v18
        case "foo bar":
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint8 = 120
                var v21 int = 175
                _ = len(v1)
                _ = fi19
                _ = v20
                _ = v21
            }
        case "\n\t":
            v3[v4] = 2147483647
        }
        _ = v17
    }
    if (-1 == -32768) {
        for fi22 := 0; fi22 < 7; fi22++ {
            v4 = (("hello" + v4) + v4)
            fmt.Println(fi22, fi22)
            _ = fi22
        }
        v23 := GenS3{F0: 4294967295}
        var v24 uint32 = 42
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := v1[1]
            _ = fi25
            _ = v26
        }
        v4 = ""
        v4 = v4
    }
    if true {
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := GenS4{F0: (("\n\t" + v4) + ((v4 + ((v4 + (v4 + v4)) + "\n\t")) + v4)), F1: 127, F2: true}
            _ = fi27
            _ = v28
        }
    } else {
        v29 := GenS1{F0: 42, F1: (!true), F2: true}
        if (false || true) {
            var v30 bool = false
            for fi31 := 0; fi31 < 5; fi31++ {
                v4 = v4
                _ = fi31
            }
            _ = v30
        } else {
            var v32 int16 = -1
            _ = v32
        }
        v33 := GenS1{F0: 774, F1: true, F2: (true && false)}
        _ = v29
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 uint8 = 2
        _ = fi34
        _ = v35
    }
    if (true || (42 != 111)) {
        v4 = v4
    } else {
        v36 := GenS2{F0: 0}
        _ = v36
    }
    if (0 == 728) {
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := []int{127, fi37, (-(fi37))}
            _ = fi37
            _ = v38
        }
        v39 := GenS1{F0: 1, F1: ((74 >= 2147483647) && true), F2: (58 == 0)}
        _ = v39
    }
    var v40 int64 = 9223372036854775807
    for fi41 := 0; fi41 < 7; fi41++ {
        v4 = ""
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
