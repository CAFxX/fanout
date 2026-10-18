package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 bool; F2 uint16 }
func genFn0() string {
    var v4 int8 = 100
    var v5 float64 = ((float64(((float32(0.0) - (0.1 * 3.14159)) + ((-2.25 / float32(-2.25)) / 1.0))) + (3.14159 - (1e10 * (float64((3.14159 * 1.5)) - 0.1)))) - -2.25)
    _ = v4
    _ = v5
    return "a"
}

func main() {
    v1 := []string{""}
    var v2 float32 = (float32(((3.14159 * ((float64((((0.1 / 3.14159) - 0.1) * float32(1e10))) / float64(1e10)) + (0.1 + (-2.25 - (-2.25 - ((-2.25 * 0.1) / float64((1.5 / ((((((-2.25 * ((-2.25 * 3.14159) * 0.1)) / 1e10) * (-2.25 - (-2.25 + (0.0 * float32(0.0))))) - 0.1) * (float32(3.14159) - ((1.5 * 0.1) * 1.5))) / 1.5))))))))) - (1.5 + 0.0))) - 1e10)
    var v3 int = -7
    if (1e10 <= 3.14159) {
        switch (v3 - (v3 - (v3 ^ 2147483647))) {
        case 32768:
            var v6 uint8 = 160
            _ = v6
        default:
            v7 := GenS1{F0: 705, F1: 100, F2: false}
            _ = v7
        }
        var v8 bool = ((false || true) && ("hello" > "hello"))
        _ = v8
    }
    if false {
        var v9 int = v3
        switch (v3 - 0) {
        case 4294967295:
            for fi10 := 0; fi10 < 7; fi10++ {
                v3 += fi10
                _ = fi10
            }
        default:
            v2 = v2
        }
        _ = v9
    } else {
        fmt.Println(v2, v2)
    }
    if true {
        v2 = 1.5
    } else {
        v11 := GenS1{F0: 181, F1: -128, F2: false}
        var v12 uint32 = 2
        _ = v11
        _ = v12
    }
    if ("\n\t" >= ("a" + ("" + ""))) {
        if false {
            v3 -= v3
            if false {
                v3 = v3
                for fi13 := 0; fi13 < 10; fi13++ {
                    switch ("a" + "hello") {
                    case "":
                        fmt.Println(v3)
                    }
                    var v14 uint = 2
                    v15 := GenS2{F0: 1}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                fmt.Println(v3)
            } else {
                v3 = (v3 / (((v3 - (v3 - v3))) | 1))
                v3 = -9223372036854775808
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v16 := GenS3{F0: -1, F1: (18446744073709551615 >= 1), F2: 65535}
            _ = v16
        }
        v2 = ((v2 * ((0.1 - (v2 * float32(1e10))) / (1.5 + v2))) / (0.1 - 1e10))
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int = (fi17 >> ((4294967295) & 63))
            _ = fi17
            _ = v18
        }
    } else {
        v2 = 1.5
        switch 32767 {
        case 2147483647:
            if ((862 >= 42) && (42 <= 2)) {
                switch -676 {
                case -330:
                    if len(v1) > 0 {
                        v1[0] = (("" + (((("\n\t" + ((("a" + ("a" + "hello")) + "foo bar") + "\n\t")) + "foo bar") + "hello") + (("foo bar" + "foo bar") + "foo bar"))) + "a")
                    }
                default:
                    v19 := GenS3{F0: 9223372036854775807, F1: true, F2: 42}
                    _ = v19
                }
                v3 = 0
                v20 := GenS3{F0: 176, F1: false, F2: 32767}
                _ = v20
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v1 = append(v1, "foo bar")
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((v3 - (v3 * 32767)) < 32768) {
        if len(v1) > 0 {
            v1[0] = ""
        }
        var v21 int8 = -1
        _ = v21
    }
    v1 = append(v1, "foo bar")
    var v22 uint64 = 831
    var v23 int64 = 0
    v22 = (v22 << ((18446744073709551615) & 63))
    switch "foo bar" {
    case "hello":
        v22 = v22
    default:
        if (!true) {
            fmt.Println(v22, v23, v22)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
