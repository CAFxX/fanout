package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 string = ""
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    var v5 bool = (("" + (((v1 + ("\n\t" + (v1 + (v1 + "\n\t")))) + v1) + v1)) > "a")
    if false {
        var v6 int32 = 108
        v1 = (v1 + v1)
        if true {
            for fi7 := 0; fi7 < 4; fi7++ {
                if (uint8(v4) != 138) {
                    v1 = ("hello" + "hello")
                    v8 := []int{(fi7 / ((425) | 1)), 354}
                    _ = v8
                } else {
                    if false {
                        fmt.Println(v4)
                        var v9 float64 = 0.1
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 16; fi10++ {
                        var v11 int16 = 770
                        if (!(v5 || ((v11 / (((v11 + v11)) | 1)) >= (v11 % ((v11) | 1))))) {
                            fmt.Println(v4, fi7)
                            v2 = (v2 / ((v2) | 1))
                            v6 += -122
                        }
                        _ = fi10
                        _ = v11
                    }
                }
                if true {
                    if (!(v1 <= "hello")) {
                        if (v5 || v5) {
                            fmt.Println(v6, v1)
                        }
                        _ = len(v3)
                        if (32767 <= fi7) {
                            var v12 uint64 = v4
                            v2 -= v2
                            fmt.Println(v6, v12)
                            _ = v12
                        }
                    } else {
                        v1 = ((((v1 + v1) + "") + (v1 + v1)) + "a")
                        var v13 uint16 = 65535
                        v3[v1] = true
                        _ = v13
                    }
                }
                if ((v1 < v1) || (!v5)) {
                    v4 = (v4 >> ((4294967295) & 63))
                    if true {
                        v14 := GenS1{F0: -2147483648, F1: (^(v4)), F2: 1}
                        _ = v14
                    } else {
                        v5 = (0 >= 42)
                    }
                    v15 := GenS1{F0: int32(v4), F1: v4, F2: 65535}
                    _ = v15
                }
                _ = fi7
            }
        }
        _ = v6
    }
    v2 = 810
    v16 := GenS2{F0: 2147483647}
    v17 := []int{-9223372036854775808}
    if (!(v1 != v1)) {
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 bool = true
            switch "" {
            case "hello":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            _ = fi18
            _ = v19
        }
    } else {
        v1 = v1
        var v20 uint64 = (v4 & (v4 << ((1) & 63)))
        _ = v20
    }
    var v21 float32 = 0.0
    v2 += v2
    v22 := GenS3{F0: v2}
    var v23 float64 = float64(v21)
    fmt.Println(v21)
    v24 := GenS2{F0: 128}
    var v25 uint64 = (v4 & (v4 + 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
