package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int {
    var v5 uint = 154
    v5 += (v5 & (v5 / ((v5) | 1)))
    _ = v5
    return -123
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 bool = (!false)
    fmt.Println(v1, v4)
    v4 = v4
    var v6 int32 = 0
    v4 = false
    for fi7 := 0; fi7 < 10; fi7++ {
        if (!((v4 || (!true)) || ((false && (2 > 65535)) && v4))) {
            if (!v4) {
                switch (("foo bar" + ((("" + "") + ("" + ("hello" + "a"))) + "a")) + "hello") {
                case "":
                    switch fi7 {
                    case 771:
                        if ((^(fi7)) <= genFn0()) {
                            v1 -= (v1 + v1)
                            var v8 int64 = (v1 * 44)
                            if (v4 || ((1e10 * (0.1 - 3.14159)) > 1.5)) {
                                var v9 string = "foo bar"
                                _ = v9
                            } else {
                                delete(v3, genFn0())
                                v6 -= 0
                            }
                            _ = v8
                        }
                    }
                case "hello":
                    fmt.Println(fi7)
                case "\n\t":
                    v10 := GenS1{F0: -217}
                    _ = v10
                }
                v1 = 42
            } else {
                var v11 uint16 = 255
                var v12 int8 = -128
                var v13 uint16 = v11
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 int = 790
            _ = v14
        } else {
            var v15 uint16 = 65535
            v4 = v4
            _ = v15
        }
        _ = fi7
    }
    v16 := GenS1{F0: -32768}
    var v17 bool = false
    if (false || v4) {
        fmt.Println(v4, v4)
        delete(v2, "")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
