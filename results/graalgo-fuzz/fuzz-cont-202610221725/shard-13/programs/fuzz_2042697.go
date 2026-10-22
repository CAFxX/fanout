package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int }
func genFn0(p0 uint32, p1 uint, p2 int32) uint16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            p0 = uint32(p1)
        case "hello":
            fmt.Println(fi5)
        case "a":
            var v6 uint32 = (p0 << (((p1 & p1)) & 63))
            _ = v6
        }
        v7 := GenS1{F0: -778, F1: (-(p2)), F2: -31}
        if false {
            var v8 bool = false
            var v9 int16 = 1
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v7
    }
    var v10 int = 548
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 255
}
func genFn1(p0 int32) int32 {
    var v11 uint = 230
    switch -101 {
    case 32767:
        var v12 uint32 = 0
        _ = v12
    case 1:
        v13 := GenS1{F0: -9223372036854775808, F1: p0, F2: 0}
        _ = v13
    default:
        var v14 int16 = -32768
        _ = v14
    }
    _ = p0
    _ = v11
    return p0
}

func main() {
    v1 := []string{"a", "hello"}
    v2 := []string{"hello", "foo bar"}
    v3 := []uint32{2147483648, 42, 4294967295}
    var v4 int8 = 75
    for fi15 := 0; fi15 < 3; fi15++ {
        if (32767 != -901) {
            if true {
                v4 += 100
                fmt.Println(v4, fi15)
                var v16 int8 = -80
                _ = v16
            }
            var v17 int16 = 32767
            _ = v17
        }
        _ = len(v2)
        _ = fi15
    }
    switch "hello" {
    case "a":
        fmt.Println(v4)
    default:
        if true {
            switch -148 {
            case 32767:
                fmt.Println(v4)
            }
            var v18 uint32 = 2147483648
            _ = v18
        }
    }
    if ((2147483647 != -63) && (!(87 < 964))) {
        var v19 float32 = 3.14159
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint8 = 0
        v22 := GenS2{F0: 1, F1: fi20}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 bool = ((!(true && (((0.0 / 1.0) * 0.0) == float32(-2.25)))) && false)
    fmt.Println(v4, v4)
    v24 := GenS3{F0: 0}
    v23 = (!(0.1 <= -2.25))
    v23 = v23
    var v25 float64 = -2.25
    v26 := GenS1{F0: 36, F1: 2147483647, F2: v4}
    switch -983 {
    case 42:
        switch "a" {
        case "a":
            var v27 float32 = -2.25
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 18; fi28++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v3 = append(v3, 987)
                _ = fi28
            }
        case "hello":
            fmt.Println(v4, v25, v25)
        default:
            v25 = 1e10
        }
    case 857:
        v23 = ((4294967295 < 2) || (false && false))
    default:
        var v29 uint32 = 42
        _ = v29
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 int32 = -1
        for fi32 := 0; fi32 < 12; fi32++ {
            if (v23 || (!(0 >= 251))) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                v33 := GenS2{F0: (v31 / ((v31) | 1)), F1: fi32}
                switch "hello" {
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    var v34 int8 = v4
                    var v35 float64 = -2.25
                    switch (fi32 ^ (-(fi32))) {
                    case 42:
                        var v36 uint = 42
                        var v37 float64 = v25
                        var v38 string = ("foo bar" + "")
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = v34
                    _ = v35
                }
                _ = v33
            } else {
                v1 = append(v1, ((("\n\t" + "\n\t") + (((("\n\t" + "") + "hello") + "hello") + ("" + (("\n\t" + "hello") + "a")))) + (("hello" + ("" + "")) + "\n\t")))
            }
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v4, v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
