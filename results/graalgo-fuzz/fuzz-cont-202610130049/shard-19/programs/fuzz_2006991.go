package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0(p0 bool) string {
    var v4 string = ""
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        if p0 {
            if true {
                if p0 {
                    var v6 int16 = 950
                    v7 := GenS1{F0: (v6 >> ((65535) & 63))}
                    v8 := []int{(fi5 - (fi5 + 895))}
                    _ = v6
                    _ = v7
                    _ = v8
                }
                var v9 string = v4
                _ = v9
            }
            v10 := []int{(fi5 & fi5)}
            var v11 int64 = 0
            _ = v10
            _ = v11
        }
        var v12 int = (fi5 * (fi5 - (fi5 << ((1) & 63))))
        _ = fi5
        _ = v12
    }
    v13 := []int{}
    _ = p0
    _ = v13
    _ = v4
    return "foo bar"
}

func main() {
    var v1 string = ("hello" + ("" + "\n\t"))
    var v2 int = 0
    var v3 float64 = 3.14159
    v3 = (v3 + v3)
    fmt.Println(v1, v1, v3)
    v14 := []int{0, (v2 << ((2) & 63)), 9223372036854775807, (-(v2))}
    v15 := GenS2{F0: 0, F1: 2, F2: 100}
    if false {
        var v16 uint32 = 65535
        _ = v16
    } else {
        var v17 uint64 = 42
        var v18 int = (v2 << ((42) & 63))
        _ = v17
        _ = v18
    }
    v1 = (("foo bar" + "\n\t") + "\n\t")
    var v19 uint8 = 42
    v3 = ((v3 + 0.0) * (float64(1.5) / 1e10))
    if false {
        var v20 uint8 = (^(v19))
        v14 = append(v14, v2)
        _ = v20
    } else {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        if true {
            var v21 uint64 = 0
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v21
        } else {
            v22 := GenS2{F0: v19, F1: (v19 * (v19 / ((v19) | 1))), F2: 1}
            v19 = (v19 + v19)
            if (!false) {
                for fi23 := 0; fi23 < 17; fi23++ {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 int8 = -1
                        v1 = ""
                        _ = fi24
                        _ = v25
                    }
                    if (true || true) {
                        var v26 float64 = (v3 - v3)
                        v27 := GenS3{F0: (v19 % ((1) | 1)), F1: 18446744073709551615, F2: v19}
                        if len(v14) > 0 {
                            v14[0] = v2
                        }
                        _ = v26
                        _ = v27
                    }
                    switch v1 {
                    case "":
                        fmt.Println(v2, v19)
                    case "hello":
                        v3 = ((v3 + v3) * (0.0 * v3))
                    }
                    _ = fi23
                }
                fmt.Println(v3)
            }
            _ = v22
        }
    }
    switch (("a" + v1) + ("hello" + "foo bar")) {
    case "hello":
        if false {
            fmt.Println(v2, v19)
            v14 = append(v14, (v2 - v2))
            v2 = v2
        }
    }
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
}
