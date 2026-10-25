package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64) string {
    var v4 string = "a"
    v5 := []int{128, 1, 0, 9223372036854775807}
    _ = p0
    _ = v4
    _ = v5
    return (v4 + (("foo bar" + "hello") + v4))
}

func main() {
    var v1 float32 = float32(0.1)
    v2 := []string{""}
    var v3 float32 = v1
    var v6 int8 = 8
    v7 := GenS1{F0: "a", F1: (v6 % (((v6 - (v6 / ((v6) | 1)))) | 1)), F2: v6}
    v6 -= v6
    fmt.Println(v6)
    if false {
        if false {
            switch -2147483648 {
            case 0:
                switch (("hello" + "") + "\n\t") {
                case "foo bar":
                    _ = len(v2)
                case "":
                    var v8 uint16 = 372
                    _ = v8
                case "a":
                    v9 := GenS1{F0: "", F1: (v6 + (v6 + 100)), F2: v6}
                    _ = v9
                }
            case 128:
                for fi10 := 0; fi10 < 11; fi10++ {
                    v2 = append(v2, ("hello" + ("foo bar" + "")))
                    switch ("\n\t" + "") {
                    case "hello":
                        v11 := GenS1{F0: ("a" + "\n\t"), F1: v6, F2: 1}
                        _ = v11
                    case "":
                        var v12 string = "hello"
                        _ = v12
                    case "\n\t":
                        var v13 uint8 = 222
                        _ = v13
                    }
                    _ = fi10
                }
            default:
                if (true && true) {
                    var v14 int = 495
                    _ = v14
                } else {
                    v6 = (v6 & (v6 / (((v6 ^ (v6 % ((v6) | 1)))) | 1)))
                    switch 32768 {
                    case -417:
                        for fi15 := 0; fi15 < 16; fi15++ {
                            v16 := GenS1{F0: "foo bar", F1: (v6 - (v6 * (v6 | (v6 + (v6 & v6))))), F2: v6}
                            var v17 uint64 = 4294967295
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    default:
                        v3 = 3.14159
                    }
                }
            }
        } else {
            var v18 int32 = -2147483648
            v1 += float32(0.1)
            _ = v18
        }
    } else {
        v3 = 0.0
    }
    var v19 bool = (!(0 < 127))
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println(v6, v6)
    v20 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
}
