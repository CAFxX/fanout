package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32) uint8 {
    v6 := []int{-118, 9223372036854775807}
    for fi7 := 0; fi7 < 10; fi7++ {
        if len(v6) > 0 {
            v6[0] = fi7
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v8 uint32 = (p0 * p0)
        _ = fi7
        _ = v8
    }
    p0 += 0
    _ = p0
    _ = v6
    return uint8(p0)
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 uint16 = 1
    var v4 string = ""
    var v5 int = int(v3)
    v5 = (v5 >> ((82) & 63))
    if ((float64(-2.25) <= 0.0) || (-1 <= -84)) {
        fmt.Println(v4, v5, v5)
        v9 := GenS1{F0: (^(v5)), F1: 65535}
        v10 := GenS1{F0: v5, F1: 571}
        _ = v10
        _ = v9
    } else {
        var v11 uint = 65535
        v5 += -1
        var v12 uint64 = 4294967295
        _ = v11
        _ = v12
    }
    v4 = "\n\t"
    if true {
        fmt.Println(v5, v3, v3)
        fmt.Println(v5, v3, v5)
        v13 := []int{}
        _ = v13
    } else {
        var v14 int8 = 0
        var v15 uint32 = 65535
        v4 = "foo bar"
        _ = v14
        _ = v15
    }
    v3 = v3
    var v16 bool = false
    v17 := []int{42}
    v3 -= 42
    v3 = 65535
    if (2147483648 == 4294967295) {
        var v18 uint64 = 18446744073709551615
        _ = v18
    } else {
        var v19 int = 4294967295
        delete(v1, v4)
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = v19
    }
    v20 := GenS2{F0: 25, F1: 127}
    var v21 int16 = 42
    if true {
        for fi22 := 0; fi22 < 8; fi22++ {
            v5 = (v5 - -47)
            switch int(v3) {
            case 1:
                switch fi22 {
                case 2147483647:
                    switch (("\n\t" + v4) + v4) {
                    case "foo bar":
                        v23 := GenS3{F0: 0}
                        _ = v23
                    case "":
                        var v24 int32 = -2147483648
                        _ = v24
                    case "\n\t":
                        delete(v1, v4)
                    }
                case 1:
                    if v16 {
                        v25 := v1[((v4 + v4) + v4)]
                        fmt.Println(v25, fi22)
                        _ = v25
                    }
                case 32768:
                    fmt.Println(v5)
                }
            case -822:
                v26 := GenS1{F0: fi22, F1: 531}
                _ = v26
            case -678:
                v27 := GenS4{F0: 2147483647}
                _ = v27
            }
            _ = fi22
        }
        v16 = v16
    } else {
        v5 += (v5 << ((0) & 63))
        var v28 uint32 = 214
        var v29 uint = 696
        _ = v28
        _ = v29
    }
    var v30 uint = 18446744073709551615
    v31 := GenS2{F0: 42, F1: 2}
    var v32 string = ((v4 + v4) + "foo bar")
    if len(v17) > 0 {
        v17[0] = v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
