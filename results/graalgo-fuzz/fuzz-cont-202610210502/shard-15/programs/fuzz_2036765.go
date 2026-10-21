package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64 }
func genFn0(p0 uint64) uint32 {
    var v4 float64 = 1.5
    if true {
        v5 := []int{127, -915, -2147483648, 32768}
        _ = v5
    } else {
        v6 := GenS1{F0: true, F1: -389, F2: p0}
        v4 = 1.5
        _ = v6
    }
    v7 := []int{}
    for fi8 := 0; fi8 < 5; fi8++ {
        p0 = (p0 + (^(p0)))
        _ = fi8
    }
    _ = p0
    _ = v4
    _ = v7
    return 65535
}

func main() {
    var v1 string = "a"
    var v2 float32 = 3.14159
    var v3 uint64 = 2
    v3 = 0
    var v9 int8 = 25
    v9 += (v9 ^ v9)
    switch v1 {
    case "foo bar":
        if (v9 == v9) {
            for fi10 := 0; fi10 < 10; fi10++ {
                switch v1 {
                case "a":
                    v1 = v1
                case "\n\t":
                    v11 := []int{(-(fi10)), 2147483647}
                    _ = v11
                case "hello":
                    v12 := GenS1{F0: false, F1: -84, F2: 0}
                    _ = v12
                }
                v13 := GenS1{F0: true, F1: 1, F2: (v3 ^ 4294967295)}
                var v14 bool = (!false)
                _ = fi10
                _ = v13
                _ = v14
            }
        } else {
            fmt.Println(v9, v2)
            fmt.Println(v1, v1, v3)
        }
    default:
        if (!((v9 << ((4294967295) & 63)) > (-(v9)))) {
            var v15 uint16 = 0
            if false {
                v15 -= v15
            } else {
                fmt.Println(v9)
            }
            var v16 float32 = 1e10
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v1, v1)
    var v17 int = -352
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            switch v1 {
            case "\n\t":
                var v20 uint16 = 4
                _ = v20
            case "a":
                if false {
                    v3 = (^(v3))
                    fmt.Println(fi19, fi18)
                }
            }
            v21 := GenS2{F0: ((v9 | (v9 >> ((621) & 63))) < v9)}
            switch v1 {
            case "foo bar":
                v22 := GenS3{F0: (fi18 | fi18), F1: -9223372036854775808}
                _ = v22
            case "a":
                v23 := []int{(fi18 * (fi19 ^ v17)), 2147483647}
                _ = v23
            }
            _ = fi19
            _ = v21
        }
        var v24 uint16 = 1
        _ = fi18
        _ = v24
    }
    var v25 uint32 = 397
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := []int{int(v25)}
        v28 := GenS1{F0: (!(!true)), F1: -1, F2: (v3 | (^(v3)))}
        fmt.Println(v3, v1, v25)
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint8 = 128
    var v30 float32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v9
}
