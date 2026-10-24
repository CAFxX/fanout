package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 int, p1 int, p2 uint) uint {
    p1 = p0
    if false {
        fmt.Println(p1, p0, p0)
        var v4 bool = ((0.1 * 0.0) >= float32(0.0))
        _ = v4
    } else {
        p1 = (p1 << ((p2) & 63))
        for fi5 := 0; fi5 < 6; fi5++ {
            p2 = p2
            p1 = (p0 ^ p1)
            _ = fi5
        }
    }
    var v6 uint = p2
    if true {
        fmt.Println(p0, p0)
    }
    switch (p0 ^ -1) {
    case 42:
        switch p0 {
        case 9223372036854775807:
            if (!(!true)) {
                v6 = 65535
                var v7 uint = p2
                _ = v7
            }
        case 630:
            v8 := GenS1{F0: (p2 * (p2 + 239))}
            _ = v8
        }
    case 32768:
        v6 -= v6
    default:
        v9 := []int{p0, p1}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return v6
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 1
    var v3 string = (("" + ("foo bar" + ("hello" + ("hello" + ("" + "foo bar"))))) + (("foo bar" + "") + ""))
    switch "\n\t" {
    case "":
        fmt.Println(v2, v3)
    default:
        switch -9223372036854775808 {
        case 42:
            v10 := GenS1{F0: 0}
            _ = v10
        case 907:
            var v11 int32 = 0
            _ = v11
        case 4294967295:
            var v12 bool = ("foo bar" != v3)
            _ = v12
        }
    }
    var v13 int64 = (^(v2))
    if (102 != 418) {
        v13 = v2
        fmt.Println(v2, v2)
        fmt.Println(v3)
    }
    v3 = ""
    v14 := GenS2{F0: -408}
    switch 1 {
    case 127:
        v15 := GenS3{F0: true}
        _ = v15
    case -2147483648:
        var v16 uint8 = 1
        _ = v16
    case 2147483647:
        var v17 uint8 = 1
        _ = v17
    }
    var v18 int64 = (v2 >> ((51) & 63))
    if false {
        if ((2 >= 128) && false) {
            for fi19 := 0; fi19 < 15; fi19++ {
                delete(v1, 459)
                _ = fi19
            }
            v20 := GenS4{F0: 1000}
            for fi21 := 0; fi21 < 14; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 uint = 210
                    _ = fi22
                    _ = v23
                }
                var v24 int32 = -2147483648
                _ = fi21
                _ = v24
            }
            _ = v20
        }
        v3 = v3
    }
    if (-555 == int32(v18)) {
        fmt.Println(v2, v3)
        if true {
            if ((663 == 798) && (false || true)) {
                fmt.Println(v18)
            }
            var v25 uint64 = 0
            if (!(!(255 != 255))) {
                switch (v3 + v3) {
                case "hello":
                    switch -2147483648 {
                    case -2147483648:
                        v2 = v2
                    }
                case "a":
                    v13 = (v13 | (v2 + v2))
                case "":
                    for fi26 := 0; fi26 < 18; fi26++ {
                        switch v3 {
                        case "foo bar":
                            v13 = (-(v13))
                        }
                        switch "foo bar" {
                        case "a":
                            switch (fi26 | (fi26 ^ (-(fi26)))) {
                            case 9223372036854775807:
                                switch (v3 + v3) {
                                case "":
                                    var v27 uint16 = uint16(v25)
                                    _ = v27
                                }
                            case 2147483647:
                                if (2 > 431) {
                                    if false {
                                        v18 = v13
                                    }
                                } else {
                                    v18 += (v13 << ((65535) & 63))
                                }
                            default:
                                v13 = (v13 - (^(v13)))
                            }
                        case "foo bar":
                            var v28 uint8 = 42
                            _ = v28
                        }
                        _ = fi26
                    }
                default:
                    v29 := []int{867, 1, 152}
                    _ = v29
                }
            }
            _ = v25
        }
        v1[-702] = 42
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        v18 -= v18
        var v31 int8 = 88
        _ = fi30
        _ = v31
    }
    v32 := GenS3{F0: true}
    v33 := v1[-9223372036854775808]
    if (false && false) {
        v2 -= -1
    }
    v34 := GenS1{F0: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
