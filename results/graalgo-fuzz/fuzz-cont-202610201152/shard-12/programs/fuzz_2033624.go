package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 bool }
func genFn0(p0 uint8) int32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(fi6)
        v7 := GenS1{F0: true, F1: "foo bar"}
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS2{F0: 32767, F1: 1}
            v10 := []int{fi8, (fi8 * fi8), (fi8 / (((fi6 ^ -2147483648)) | 1)), (fi6 * fi8)}
            fmt.Println(p0, fi8)
            v11 := GenS2{F0: 849, F1: 2147483648}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v12 int = -2147483648
    var v13 bool = (!((^(p0)) == (^(p0))))
    fmt.Println(v12, v13)
    _ = p0
    _ = v12
    _ = v13
    return 571
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int32 = -485
    v3 := []string{("a" + (("\n\t" + "hello") + ("a" + "hello"))), "hello", "hello"}
    var v4 bool = (98 != 4294967295)
    var v5 int64 = -1
    fmt.Println(v5, v5, v1)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println(v1, v5, v1)
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 bool = v4
        v2 = -731
        _ = fi14
        _ = v15
    }
    var v16 int16 = 1
    if false {
        if (v4 && v4) {
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 int64 = v5
                v19 := GenS3{F0: 255}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        if true {
            if ((("foo bar" + ("\n\t" + "\n\t")) + "hello") <= "foo bar") {
                v4 = v4
                fmt.Println(v4, v5)
            } else {
                if v4 {
                    var v20 uint64 = 18446744073709551615
                    _ = v20
                }
            }
        }
    }
    var v21 int32 = v2
    if (v2 <= 42) {
        v21 = (-(v2))
        var v22 int32 = genFn0(127)
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS4{F0: (-(v16)), F1: v4}
            v25 := GenS4{F0: (v16 * (-(v16))), F1: false}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    var v26 uint16 = 14
    _ = len(v3)
    if (v26 != 42) {
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS3{F0: (^(v26))}
            v1 += (-(fi27))
            _ = fi27
            _ = v28
        }
        v29 := GenS2{F0: (v16 ^ 42), F1: 65535}
        _ = v29
    } else {
        switch "\n\t" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = "a"
            }
        default:
            if (!false) {
                for fi30 := 0; fi30 < 6; fi30++ {
                    v4 = true
                    v4 = true
                    _ = fi30
                }
            }
        }
        var v31 uint16 = 32767
        _ = v31
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v21, v16, v2)
    case "":
        v21 = (-(v21))
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := GenS4{F0: v16, F1: v4}
        if (103 != 4294967295) {
            if v4 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1, v1, v1)
            } else {
                switch "hello" {
                case "":
                    v34 := GenS4{F0: (v16 ^ v16), F1: true}
                    _ = v34
                case "a":
                    fmt.Println(fi32, v2, v21)
                }
                v35 := GenS3{F0: v26}
                _ = v35
            }
            var v36 bool = v4
            var v37 uint8 = 127
            _ = v36
            _ = v37
        }
        _ = fi32
        _ = v33
    }
    var v38 string = (((("hello" + ("a" + (("\n\t" + "hello") + "a"))) + ("hello" + (("a" + "\n\t") + ("\n\t" + "")))) + (("\n\t" + "") + "")) + "a")
    fmt.Println(v26, v1)
    fmt.Println(v4, v26, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
