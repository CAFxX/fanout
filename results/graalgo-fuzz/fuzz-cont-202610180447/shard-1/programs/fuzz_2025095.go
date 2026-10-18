package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0() string {
    v6 := GenS1{F0: 2}
    var v7 int16 = -791
    var v8 int8 = -105
    v9 := []int{-39, 0, 32767, 0}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = -2.25
    var v3 uint64 = (v1 | v1)
    v4 := make(map[uint32]bool)
    v5 := []int{1, 1, -2147483648, 411}
    for fi10 := 0; fi10 < 8; fi10++ {
        if true {
            var v11 float32 = v2
            v12 := GenS1{F0: 127}
            _ = v11
            _ = v12
        }
        if len(v5) > 0 {
            v5[0] = (^(fi10))
        }
        if true {
            var v13 uint16 = 862
            if false {
                _ = len(v5)
                v13 += (v13 / ((42) | 1))
                switch "\n\t" {
                case "foo bar":
                    switch "\n\t" {
                    case "a":
                        switch "foo bar" {
                        case "hello":
                            v3 += v1
                        case "":
                            var v14 float64 = 0.1
                            _ = v14
                        case "\n\t":
                            v3 = v3
                        }
                    case "foo bar":
                        var v15 int64 = int64(fi10)
                        _ = v15
                    }
                case "a":
                    var v16 float32 = (((v2 + 1e10) * float32((1.5 - 1e10))) / (1.5 + ((((v2 + v2) * 0.0) + 1.5) * 0.0)))
                    _ = v16
                case "hello":
                    v2 += (0.1 - (v2 + 1e10))
                }
            } else {
                fmt.Println(v13, fi10)
                for fi17 := 0; fi17 < 4; fi17++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v18 := GenS1{F0: 255}
                    _ = fi17
                    _ = v18
                }
            }
            var v19 int16 = 1
            _ = v13
            _ = v19
        } else {
            var v20 int64 = 42
            _ = v20
        }
        _ = fi10
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        fmt.Println(v1, v3)
        _ = fi21
    }
    var v22 int = int(v1)
    switch (("\n\t" + ((("foo bar" + "foo bar") + "a") + "\n\t")) + ("\n\t" + "foo bar")) {
    case "foo bar":
        v23 := GenS2{F0: -103, F1: 32767}
        _ = v23
    }
    v5 = append(v5, v22)
    if true {
        for fi24 := 0; fi24 < 14; fi24++ {
            if (!true) {
                v25 := GenS1{F0: 255}
                v4[0] = false
                _ = v25
            } else {
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 bool = (!(((true && true) || true) || false))
                    var v28 uint32 = 4294967295
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                if true {
                    var v29 uint16 = 747
                    v29 = v29
                    _ = v29
                }
            }
            _ = fi24
        }
        v22 += v22
        var v30 bool = false
        _ = v30
    }
    var v31 uint = 107
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 int16 = 32767
        _ = fi32
        _ = v33
    }
    var v34 int32 = 42
    v22 -= 4294967295
    var v35 bool = true
    var v36 int32 = v34
    var v37 bool = v35
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
