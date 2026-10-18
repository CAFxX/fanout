package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 bool, p1 int, p2 string) uint8 {
    fmt.Println(p2)
    v7 := GenS1{F0: 651}
    var v8 int8 = 127
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = 100
    var v2 int16 = 32767
    var v3 string = "hello"
    var v4 uint64 = 18446744073709551615
    var v5 int = -2147483648
    var v6 string = v3
    v9 := GenS2{F0: -950, F1: 54}
    v10 := GenS2{F0: 2147483647, F1: 2}
    switch "foo bar" {
    case "\n\t":
        if true {
            var v11 int8 = -1
            _ = v11
        }
    default:
        switch v6 {
        case "foo bar":
            var v12 int8 = (v1 << ((255) & 63))
            _ = v12
        case "\n\t":
            switch v3 {
            case "a":
                v13 := []int{(v5 * -322), (v5 >> ((0) & 63))}
                _ = v13
            case "hello":
                if true {
                    var v14 uint8 = 1
                    v15 := []int{}
                    _ = v14
                    _ = v15
                }
            default:
                var v16 int = 128
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v2 = v2
        var v18 bool = true
        var v19 uint32 = 489
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS1{F0: (v4 + v4)}
        var v22 int = fi20
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := []int{(-(v5)), v22}
            var v25 int32 = 520
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    v3 = ((v3 + (v6 + v3)) + v6)
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := []int{fi26, fi26, v5}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS3{F0: 0}
        if true {
            v4 = (v4 / ((v4) | 1))
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 uint16 = 48
                _ = fi30
                _ = v31
            }
        } else {
            var v32 uint64 = 2
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    var v33 float32 = (1e10 - 3.14159)
    if true {
        for fi34 := 0; fi34 < 15; fi34++ {
            var v35 uint8 = 223
            var v36 bool = false
            var v37 int16 = (v2 % ((232) | 1))
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
    }
    if (!("\n\t" < ("" + "a"))) {
        v1 -= v1
        v38 := []int{(v5 % (((-(v5))) | 1))}
        _ = v38
    }
    if false {
        if (!false) {
            fmt.Println(v5, v2)
            for fi39 := 0; fi39 < 17; fi39++ {
                var v40 uint32 = 2147483647
                v40 -= 4294967295
                _ = fi39
                _ = v40
            }
            var v41 int8 = -128
            _ = v41
        } else {
            switch (v5 >> ((0) & 63)) {
            case 32768:
                var v42 uint64 = v4
                _ = v42
            case -1:
                v43 := GenS4{F0: 163, F1: "hello"}
                _ = v43
            case -2147483648:
                var v44 uint8 = 127
                _ = v44
            }
            var v45 bool = (true && true)
            _ = v45
        }
    } else {
        for fi46 := 0; fi46 < 10; fi46++ {
            var v47 int64 = 737
            _ = fi46
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
