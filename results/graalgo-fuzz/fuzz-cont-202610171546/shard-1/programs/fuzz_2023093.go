package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint, p1 string) int32 {
    if (42 != -9223372036854775808) {
        var v4 int32 = 0
        _ = v4
    }
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{fi5, fi5}
        if (65535 == 2) {
            fmt.Println(fi5, p0, fi5)
            var v7 int = (^(fi5))
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int32 = 42
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    var v10 uint16 = 131
    v11 := GenS1{F0: 128, F1: p1, F2: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := []string{("\n\t" + ("hello" + (("hello" + "hello") + ((("foo bar" + "") + "foo bar") + "\n\t")))), (("hello" + "a") + "foo bar"), "a"}
    var v2 uint32 = 323
    v3 := []int{-370, 4294967295, 32767}
    var v12 int16 = -194
    if false {
        if true {
            var v13 uint64 = 18446744073709551615
            fmt.Println(v2)
            var v14 bool = false
            _ = v13
            _ = v14
        }
        v2 -= v2
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int64 = 42
        v3 = append(v3, fi15)
        v2 = (v2 & (v2 << ((1) & 63)))
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 2}
    v18 := GenS2{F0: 32767}
    v19 := GenS2{F0: 42}
    v20 := GenS3{F0: 814, F1: 42, F2: 42}
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(v12, fi21)
        _ = fi21
    }
    var v22 string = ("" + "\n\t")
    _ = len(v3)
    v23 := GenS2{F0: 255}
    if false {
        v22 = "hello"
        if (!true) {
            v2 = v2
            v24 := GenS4{F0: -35, F1: 0}
            _ = v24
        }
        v25 := GenS3{F0: 275, F1: 0, F2: 786}
        _ = v25
    } else {
        v26 := GenS4{F0: 9223372036854775807, F1: 0}
        v12 = (v12 | v12)
        _ = v26
    }
    _ = len(v1)
    fmt.Println(v22, v22, v22)
    switch 0 {
    case 9223372036854775807:
        var v27 string = (((v22 + ((v22 + v22) + "foo bar")) + v22) + (v22 + v22))
        _ = v27
    case 128:
        switch "\n\t" {
        case "a":
            v2 = (v2 | 4294967295)
        case "":
            var v28 uint8 = 2
            _ = v28
        default:
            if true {
                v29 := GenS4{F0: -1, F1: 127}
                fmt.Println(v12)
                fmt.Println(v22, v22)
                _ = v29
            } else {
                for fi30 := 0; fi30 < 20; fi30++ {
                    v31 := GenS2{F0: 597}
                    v32 := GenS2{F0: 65535}
                    v22 = v22
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        }
    case 32768:
        v12 += (v12 - -394)
    }
    v22 = "a"
    if (-147 == 4294967295) {
        var v33 uint = 18446744073709551615
        fmt.Println(v12, v33)
        for fi34 := 0; fi34 < 7; fi34++ {
            v33 = 255
            _ = fi34
        }
        _ = v33
    } else {
        switch v22 {
        case "foo bar":
            for fi35 := 0; fi35 < 18; fi35++ {
                for fi36 := 0; fi36 < 4; fi36++ {
                    fmt.Println(v22, fi35)
                    v37 := GenS5{F0: 854}
                    _ = fi36
                    _ = v37
                }
                v1 = append(v1, v22)
                _ = fi35
            }
        case "":
            v12 -= (v12 >> ((42) & 63))
        case "hello":
            switch -1 {
            case 0:
                v38 := GenS3{F0: 18446744073709551615, F1: 65535, F2: 0}
                _ = v38
            }
        default:
            if (!true) {
                var v39 int16 = (v12 >> ((4294967295) & 63))
                v22 = ((v22 + "a") + "hello")
                _ = v39
            }
        }
        v2 += v2
        var v40 int = 42
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
}
