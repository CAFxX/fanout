package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0(p0 int16, p1 int16) uint16 {
    if (65535 == 365) {
        var v4 int16 = p0
        _ = v4
    }
    var v5 bool = true
    fmt.Println(p0)
    v6 := []int{0, -259}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float32 = -2.25
    var v3 int = 4294967295
    fmt.Println(v1, v2)
    fmt.Println(v2, v2, v1)
    for fi7 := 0; fi7 < 4; fi7++ {
        switch "a" {
        case "a":
            if false {
                var v8 uint = 1
                v9 := GenS1{F0: "a", F1: 764, F2: (v1 >> (((^(v8))) & 63))}
                _ = v8
                _ = v9
            }
        default:
            if true {
                var v10 int32 = 42
                _ = v10
            }
        }
        var v11 float32 = 1.5
        var v12 uint32 = 44
        _ = fi7
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS1{F0: ("" + ""), F1: -2147483648, F2: -837}
        _ = fi13
        _ = v14
    }
    v3 = v3
    v3 -= v3
    v15 := GenS2{F0: 4294967295, F1: 4294967295}
    var v16 string = "a"
    switch (-(v3)) {
    case 2147483647:
        if true {
            var v17 uint64 = 2
            v3 -= 609
            var v18 int8 = 127
            _ = v17
            _ = v18
        }
    case 42:
        v1 -= 42
    default:
        if false {
            v19 := GenS3{F0: 428}
            for fi20 := 0; fi20 < 10; fi20++ {
                switch v3 {
                case 4294967295:
                    var v21 int = -1
                    _ = v21
                default:
                    var v22 uint32 = 65535
                    _ = v22
                }
                v23 := []int{fi20, fi20, (^(fi20)), (^(fi20))}
                var v24 int16 = 32767
                _ = fi20
                _ = v23
                _ = v24
            }
            _ = v19
        }
    }
    v25 := GenS3{F0: 2147483648}
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint8 = 1
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := GenS1{F0: (v16 + v16), F1: -1, F2: v1}
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    if (0 < 989) {
        v30 := GenS3{F0: 2147483648}
        for fi31 := 0; fi31 < 12; fi31++ {
            v32 := []int{(-(fi31)), fi31}
            if true {
                v33 := GenS4{F0: -844, F1: 65535, F2: -2147483648}
                _ = v33
            }
            for fi34 := 0; fi34 < 4; fi34++ {
                switch (fi31 / ((fi31) | 1)) {
                case 9223372036854775807:
                    switch v16 {
                    case "":
                        v2 = v2
                    case "foo bar":
                        if false {
                            v16 = v16
                            v3 = (^(fi34))
                        }
                    default:
                        if (536 >= 165) {
                            switch v16 {
                            case "a":
                                if (true || (v2 != float32(0.0))) {
                                    var v35 uint = 4294967295
                                    _ = v35
                                }
                            case "foo bar":
                                fmt.Println(v2)
                            }
                        }
                    }
                case 32768:
                    v16 = ""
                    v16 = v16
                    fmt.Println(fi31)
                default:
                    switch (v3 ^ 616) {
                    case 127:
                        var v36 uint8 = 42
                        _ = v36
                    }
                }
                v37 := GenS4{F0: 2147483647, F1: 0, F2: 0}
                _ = fi34
                _ = v37
            }
            _ = fi31
            _ = v32
        }
        if false {
            var v38 int64 = v1
            fmt.Println(v16, v16, v3)
            switch v16 {
            case "":
                for fi39 := 0; fi39 < 4; fi39++ {
                    var v40 int = fi39
                    _ = fi39
                    _ = v40
                }
            }
            _ = v38
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
