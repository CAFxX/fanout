package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int {
    var v5 string = ""
    v5 = "\n\t"
    var v6 float64 = 1.5
    var v7 int32 = -731
    if false {
        v6 = 1e10
        for fi8 := 0; fi8 < 3; fi8++ {
            v5 = (v5 + ("\n\t" + (v5 + "a")))
            _ = fi8
        }
        v9 := []int{}
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 391
    var v2 uint32 = 2147483647
    v3 := []string{"hello"}
    var v4 int = 2147483647
    if (!true) {
        if false {
            _ = len(v3)
            v10 := GenS1{F0: uint(v1)}
            _ = v10
        }
        fmt.Println(v1)
    }
    if (false || (-2.25 > 1e10)) {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(fi11)
                var v13 int32 = -213
                _ = fi12
                _ = v13
            }
            var v14 float64 = (0.1 * 0.1)
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v15 uint8 = 2
                for fi16 := 0; fi16 < 7; fi16++ {
                    v14 = v14
                    var v17 int64 = 9223372036854775807
                    if (false || true) {
                        switch "" {
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    v18 := GenS1{F0: 255}
                    var v19 int8 = 1
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = v15
            } else {
                var v20 float32 = 0.0
                v4 += (v4 - v4)
                switch (v4 | v4) {
                case 1:
                    fmt.Println(v14, v4)
                case 128:
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + "hello")
                    }
                }
                _ = v20
            }
            _ = fi11
            _ = v14
        }
        v21 := GenS1{F0: 255}
        _ = v21
    }
    v2 = (v2 + 2147483647)
    v22 := GenS1{F0: 1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4, v4, v2)
    if (((v1 & 2147483647) > (^(v2))) && true) {
        var v23 uint16 = 40
        _ = v23
    } else {
        var v24 int64 = 9223372036854775807
        v25 := GenS1{F0: 1}
        var v26 int16 = 32767
        _ = v24
        _ = v25
        _ = v26
    }
    var v27 bool = false
    if (v27 || true) {
        fmt.Println(v2, v4, v1)
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 int16 = -959
            _ = fi28
            _ = v29
        }
    }
    fmt.Println(v4)
    fmt.Println(v27, v4, v2)
    switch v4 {
    case 127:
        var v30 int32 = 2147483647
        _ = v30
    }
    var v31 float32 = float32(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
