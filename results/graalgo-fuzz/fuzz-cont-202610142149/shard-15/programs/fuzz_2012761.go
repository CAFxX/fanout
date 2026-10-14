package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 int64 }

func main() {
    var v1 uint64 = 2
    var v2 uint32 = 2
    var v3 uint16 = 737
    var v4 bool = (!true)
    v5 := make(map[int]uint32)
    v6 := make(map[string]bool)
    if v4 {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        switch -2147483648 {
        case -677:
            var v9 string = "foo bar"
            _ = v9
        default:
            v2 = (v2 % (((v2 ^ v2)) | 1))
        }
        v10 := []int{4294967295, -1}
        _ = v10
    } else {
        switch ("hello" + "a") {
        case "a":
            v3 += v3
        case "\n\t":
            v11 := GenS1{F0: true, F1: 661}
            _ = v11
        }
    }
    v4 = (!v4)
    if (132 <= 13) {
        for fi12 := 0; fi12 < 17; fi12++ {
            if (32767 <= -1) {
                var v13 int64 = 0
                _ = v13
            }
            v14 := GenS1{F0: ((-6 == -68) || (1 == 255)), F1: v1}
            _ = fi12
            _ = v14
        }
        switch 127 {
        case -662:
            switch 127 {
            case -323:
                v2 = v2
            }
        }
    } else {
        switch -9223372036854775808 {
        case 2147483647:
            delete(v6, "\n\t")
        case -585:
            v15 := GenS1{F0: true, F1: (v1 / (((v1 % (((^(v1))) | 1))) | 1))}
            _ = v15
        default:
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int8 = 100
                v18 := GenS1{F0: (222 <= 189), F1: (^(v1))}
                var v19 int64 = -437
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        }
    }
    fmt.Println(v4)
    var v20 int64 = -443
    var v21 int32 = 42
    if false {
        if v4 {
            v22 := []int{-1, int(v1), 1, -507}
            _ = v22
        } else {
            v23 := GenS2{F0: 0, F1: 2, F2: v20}
            v24 := []int{32767}
            if (!false) {
                if true {
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                }
                for fi25 := 0; fi25 < 12; fi25++ {
                    v20 += 1
                    _ = fi25
                }
                var v26 bool = v4
                _ = v26
            }
            _ = v23
            _ = v24
        }
        fmt.Println(v1, v3, v21)
        v21 = v21
    } else {
        v1 = v1
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v3 += (v3 >> ((298) & 63))
        v28 := v5[(fi27 << ((0) & 63))]
        var v29 int16 = 0
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int = 2147483647
    fmt.Println(v1, v4, v21)
    v20 = (-(v20))
    switch "foo bar" {
    case "a":
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 uint64 = v1
            fmt.Println(v2)
            _ = fi31
            _ = v32
        }
    default:
        fmt.Println(v2, v30, v3)
    }
    var v33 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
