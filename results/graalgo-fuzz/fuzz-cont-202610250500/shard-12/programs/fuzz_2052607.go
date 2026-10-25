package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
type GenS4 struct { F0 string; F1 uint32; F2 uint64 }
func genFn0(p0 uint8, p1 bool, p2 uint) int16 {
    v5 := GenS1{F0: 65535, F1: 1}
    v6 := GenS1{F0: p2, F1: 360}
    for fi7 := 0; fi7 < 6; fi7++ {
        p2 = (p2 / ((p2) | 1))
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{128, (^(fi7))}
            v10 := GenS2{F0: (^(fi7)), F1: 100}
            v9 = append(v9, -1)
            _ = fi8
            _ = v10
            _ = v9
        }
        if p1 {
            fmt.Println(fi7, p0)
        }
        _ = fi7
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    _ = v6
    return 704
}
func genFn1(p0 int64) int64 {
    p0 += p0
    for fi12 := 0; fi12 < 2; fi12++ {
        if (!(!true)) {
            for fi13 := 0; fi13 < 4; fi13++ {
                p0 = 0
                _ = fi13
            }
            for fi14 := 0; fi14 < 6; fi14++ {
                fmt.Println(p0)
                var v15 uint = 4294967295
                var v16 int32 = 2147483647
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 int16 = -32768
            _ = v17
        }
        var v18 uint16 = 2
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(fi12, fi12, fi12)
            var v20 uint8 = 204
            _ = fi19
            _ = v20
        }
        _ = fi12
        _ = v18
    }
    _ = p0
    return p0
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = 0.0
    var v3 uint = 65535
    v4 := make(map[uint32]uint32)
    v1 += (v1 * v1)
    switch "foo bar" {
    case "":
        var v21 int32 = 42
        _ = v21
    }
    var v22 int = 2147483647
    var v23 string = "\n\t"
    var v24 uint32 = 723
    var v25 uint8 = 127
    for fi26 := 0; fi26 < 8; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            if (!true) {
                v28 := GenS2{F0: (fi26 << ((957) & 63)), F1: 42}
                delete(v4, 42)
                _ = v28
            }
            v2 += v1
            fmt.Println(v1, v22)
            fmt.Println(v22, fi26)
            _ = fi27
        }
        if true {
            if (v25 == uint8(v3)) {
                v29 := GenS3{F0: (^(v3)), F1: -392, F2: int32(v22)}
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 int16 = 920
                    v32 := v4[v24]
                    var v33 uint64 = 4294967295
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                var v34 uint64 = 1
                _ = v29
                _ = v34
            }
            if false {
                var v35 int8 = 1
                _ = v35
            }
            v36 := []int{}
            _ = v36
        } else {
            switch (fi26 - v22) {
            case -517:
                if (v23 >= "foo bar") {
                    v37 := []int{fi26, 2147483647}
                    v23 = v23
                    var v38 uint64 = 2
                    _ = v37
                    _ = v38
                } else {
                    v39 := GenS3{F0: (v3 / ((v3) | 1)), F1: 4294967295, F2: 1}
                    v40 := GenS4{F0: v23, F1: v24, F2: 209}
                    var v41 float64 = ((-2.25 * 3.14159) / -2.25)
                    _ = v39
                    _ = v40
                    _ = v41
                }
            case -2147483648:
                switch (v22 ^ v22) {
                case 0:
                    v42 := []int{(fi26 << (((v3 / (((v3 >> ((42) & 63))) | 1))) & 63))}
                    _ = v42
                default:
                    v24 = (v24 + v24)
                }
            case 430:
                v2 -= v2
            }
            v22 = -500
            var v43 uint16 = 2
            _ = v43
        }
        v1 = 0.0
        _ = fi26
    }
    var v44 bool = false
    switch 127 {
    case 128:
        v45 := []int{v22, (v22 ^ v22), v22}
        _ = v45
    case 32767:
        v46 := GenS4{F0: v23, F1: (v24 << ((v3) & 63)), F2: 122}
        _ = v46
    default:
        v22 = v22
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        v48 := []int{-596}
        for fi49 := 0; fi49 < 5; fi49++ {
            var v50 float32 = v2
            var v51 bool = v44
            if v51 {
                if len(v48) >= 2 {
                    v48 = v48[1:len(v48)]
                }
                var v52 string = v23
                if len(v48) > 0 {
                    v48[0] = (^(fi47))
                }
                _ = v52
            }
            _ = fi49
            _ = v50
            _ = v51
        }
        v53 := GenS3{F0: v3, F1: (^(fi47)), F2: -643}
        _ = fi47
        _ = v48
        _ = v53
    }
    var v54 uint32 = (v24 - v24)
    v55 := v4[v54]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v54
    _ = v55
}
