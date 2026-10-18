package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 string, p1 bool) int16 {
    var v4 uint64 = 2
    switch 42 {
    case 253:
        for fi5 := 0; fi5 < 8; fi5++ {
            v6 := GenS1{F0: 66, F1: 65535}
            _ = fi5
            _ = v6
        }
    case -9223372036854775808:
        v7 := GenS1{F0: 4294967295, F1: 4294967295}
        _ = v7
    case 2147483647:
        switch p0 {
        case "hello":
            for fi8 := 0; fi8 < 4; fi8++ {
                p1 = (65535 <= 65535)
                var v9 uint64 = v4
                _ = fi8
                _ = v9
            }
        }
    default:
        if p1 {
            v10 := GenS1{F0: 1, F1: 1}
            _ = v10
        } else {
            v11 := GenS2{F0: 42}
            _ = v11
        }
    }
    v12 := []int{127, 0, 2147483647}
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 bool = p1
        if true {
            v4 -= (v4 ^ (v4 / ((18446744073709551615) | 1)))
        } else {
            v14 = p1
            p0 = ""
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 string = ("foo bar" + (p0 + "foo bar"))
                _ = fi15
                _ = v16
            }
        }
        var v17 float32 = float32(-2.25)
        _ = fi13
        _ = v14
        _ = v17
    }
    fmt.Println(p1, p1, v4)
    _ = p0
    _ = p1
    _ = v12
    _ = v4
    return 22
}
func genFn1(p0 int32, p1 int8) bool {
    var v18 int = 0
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int16 = 3
        for fi21 := 0; fi21 < 11; fi21++ {
            switch "foo bar" {
            case "foo bar":
                p1 -= p1
            }
            if false {
                v20 = 1000
            }
            v22 := GenS2{F0: 2147483647}
            _ = fi21
            _ = v22
        }
        var v23 bool = (!(fi19 > (fi19 * 127)))
        _ = fi19
        _ = v20
        _ = v23
    }
    v24 := []int{v18, 81, -1, v18}
    _ = p0
    _ = p1
    _ = v18
    _ = v24
    return false
}

func main() {
    var v1 string = "a"
    var v2 int32 = 42
    var v3 float32 = 0.1
    v3 = v3
    v25 := GenS1{F0: 2, F1: 0}
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 int64 = 42
        if (!false) {
            for fi28 := 0; fi28 < 16; fi28++ {
                switch v1 {
                case "a":
                    var v29 int32 = -1
                    _ = v29
                default:
                    v30 := GenS2{F0: 0}
                    _ = v30
                }
                var v31 int = fi26
                fmt.Println(fi26, v1)
                _ = fi28
                _ = v31
            }
            var v32 uint64 = 18446744073709551615
            v33 := GenS1{F0: 870, F1: 0}
            _ = v32
            _ = v33
        } else {
            var v34 bool = ((v1 + v1) < "")
            for fi35 := 0; fi35 < 5; fi35++ {
                v34 = (!false)
                v34 = (-32768 < genFn0(v1, (true || true)))
                var v36 float32 = float32(3.14159)
                var v37 uint = 2
                _ = fi35
                _ = v36
                _ = v37
            }
            var v38 float32 = float32(1e10)
            _ = v34
            _ = v38
        }
        var v39 uint = 197
        _ = fi26
        _ = v27
        _ = v39
    }
    var v40 uint16 = 2
    v1 = "hello"
    for fi41 := 0; fi41 < 17; fi41++ {
        for fi42 := 0; fi42 < 19; fi42++ {
            var v43 bool = ((v2 ^ (v2 / ((-876) | 1))) != v2)
            var v44 int64 = 9223372036854775807
            _ = fi42
            _ = v43
            _ = v44
        }
        switch v1 {
        case "a":
            v1 = "a"
        }
        fmt.Println(fi41, v1)
        _ = fi41
    }
    var v45 int8 = -1
    v2 = (v2 ^ (v2 << ((1) & 63)))
    v46 := GenS3{F0: (v1 + (v1 + v1)), F1: 32767, F2: 42}
    fmt.Println(v2, v45)
    var v47 uint16 = v40
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v40
    _ = v45
    _ = v46
    _ = v47
}
