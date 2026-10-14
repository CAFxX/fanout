package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16, p1 string) uint8 {
    var v7 int = 2147483647
    var v8 uint64 = 4294967295
    v9 := []int{(v7 ^ v7), 2147483647, v7}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 103
}
func genFn1(p0 uint32, p1 int32) uint64 {
    var v10 uint64 = 858
    v11 := GenS1{F0: 232}
    fmt.Println(p0, p1, v10)
    v10 += v10
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 float64 = (((1.5 / 1.0) - 0.1) / -2.25)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (^(v10))
}

func main() {
    var v1 uint64 = 889
    var v2 string = ("\n\t" + "foo bar")
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    var v5 uint16 = 32767
    v6 := []int{128}
    fmt.Println(v5, v1)
    var v14 uint8 = 1
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := GenS2{F0: 2}
        var v17 int16 = 1
        fmt.Println(v14, v14, v17)
        _ = fi15
        _ = v16
        _ = v17
    }
    v1 = v1
    _ = len(v6)
    fmt.Println(v2, v14, v1)
    for fi18 := 0; fi18 < 12; fi18++ {
        switch (fi18 - fi18) {
        case 2147483647:
            var v19 uint64 = (v1 + (v1 << ((199) & 63)))
            _ = v19
        default:
            switch (((v2 + v2) + (v2 + ((("a" + v2) + "") + "hello"))) + "\n\t") {
            case "":
                v14 += v14
            case "hello":
                v20 := GenS2{F0: 42}
                _ = v20
            }
        }
        _ = fi18
    }
    var v21 uint = 4294967295
    var v22 int32 = 1
    v23 := GenS2{F0: (^(v21))}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v24 uint64 = 825
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS3{F0: v22, F1: v14}
        switch "\n\t" {
        case "a":
            var v27 uint8 = 2
            _ = v27
        default:
            switch v2 {
            case "hello":
                for fi28 := 0; fi28 < 14; fi28++ {
                    var v29 string = v2
                    v22 = int32(v14)
                    var v30 int = (fi28 - (fi28 - fi25))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            case "\n\t":
                v14 = (v14 * (v14 / (((v14 % ((genFn0((v5 - 255), ((v2 + v2) + "hello"))) | 1))) | 1)))
            case "":
                var v31 float64 = 0.1
                _ = v31
            }
        }
        _ = fi25
        _ = v26
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 float32 = 0.0
        if false {
            if (v1 < uint64(v14)) {
                v34 := GenS3{F0: (v22 << ((v21) & 63)), F1: genFn0(v5, v2)}
                _ = v34
            } else {
                var v35 bool = true
                for fi36 := 0; fi36 < 14; fi36++ {
                    v37 := GenS4{F0: (v24 >> ((2) & 63)), F1: 42}
                    fmt.Println(v22)
                    _ = fi36
                    _ = v37
                }
                _ = v35
            }
            var v38 int8 = 0
            for fi39 := 0; fi39 < 18; fi39++ {
                v40 := GenS5{F0: 1}
                _ = fi39
                _ = v40
            }
            _ = v38
        } else {
            if true {
                var v41 int8 = 100
                _ = v41
            } else {
                if ((!true) && ((1 < 1) || (4294967295 < 490))) {
                    v21 = (v21 >> ((v21) & 63))
                }
            }
        }
        _ = fi32
        _ = v33
    }
    var v42 float32 = float32((((3.14159 - 0.0) / 0.1) / 1.0))
    if false {
        var v43 uint = (v21 << ((987) & 63))
        _ = v43
    }
    v44 := v4[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v5
    _ = v6
}
