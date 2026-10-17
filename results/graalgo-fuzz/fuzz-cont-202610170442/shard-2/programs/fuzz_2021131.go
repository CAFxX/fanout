package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0() int {
    var v7 bool = false
    if v7 {
        if v7 {
            var v8 uint8 = 42
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int32 = 0
                v11 := []int{(-(fi9)), (fi9 / (((fi9 / ((int(v10)) | 1))) | 1))}
                var v12 uint32 = 65535
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            if (-1 > -690) {
                v13 := []int{-37, 358, 522, 1}
                _ = v13
            } else {
                var v14 int8 = 0
                fmt.Println(v8)
                var v15 int16 = -1
                _ = v14
                _ = v15
            }
            _ = v8
        } else {
            v7 = ((42 > 2) && v7)
        }
    } else {
        switch "a" {
        case "\n\t":
            v16 := []int{590, -285, 2147483647, 620}
            _ = v16
        default:
            switch -2147483648 {
            case 32768:
                v7 = ((v7 || true) && true)
            }
        }
        var v17 int = 9223372036854775807
        var v18 uint32 = 4294967295
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 0, F1: -1}
    _ = v19
    _ = v7
    return 448
}

func main() {
    var v1 uint = 473
    var v2 uint8 = uint8(v1)
    var v3 int16 = -1
    v4 := make(map[int]bool)
    var v5 float64 = 1.5
    var v6 int8 = -128
    v20 := GenS1{F0: 42, F1: 1}
    for fi21 := 0; fi21 < 3; fi21++ {
        v3 -= -491
        v22 := []int{fi21, fi21}
        var v23 bool = false
        _ = fi21
        _ = v22
        _ = v23
    }
    if (2 < (^(v1))) {
        var v24 string = ""
        _ = v24
    }
    v2 += v2
    v5 -= -2.25
    for fi25 := 0; fi25 < 12; fi25++ {
        if false {
            var v26 int16 = int16(v1)
            v27 := GenS1{F0: 42, F1: -907}
            if (1 > -2147483648) {
                v28 := []int{}
                var v29 uint16 = 284
                v30 := GenS1{F0: 615, F1: -9223372036854775808}
                _ = v28
                _ = v29
                _ = v30
            } else {
                var v31 float64 = float64(-2.25)
                _ = len(v4)
                fmt.Println(v5)
                _ = v31
            }
            _ = v26
            _ = v27
        }
        if (true || false) {
            v32 := GenS2{F0: 110, F1: (v6 << ((18446744073709551615) & 63)), F2: 9223372036854775807}
            v33 := GenS2{F0: (v2 & (v2 % (((^(v2))) | 1))), F1: v6, F2: 45}
            v34 := GenS3{F0: 0, F1: 4294967295}
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi25
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS3{F0: 1, F1: 2}
        var v37 string = "hello"
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := GenS2{F0: v2, F1: v6, F2: -1}
    if true {
        var v39 uint8 = (v2 + (v2 | (v2 ^ (v2 - v2))))
        _ = v39
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        var v41 string = ("hello" + "\n\t")
        var v42 float64 = 1e10
        var v43 float32 = 3.14159
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    var v44 uint = 18446744073709551615
    for fi45 := 0; fi45 < 17; fi45++ {
        v2 = v2
        _ = fi45
    }
    var v46 int16 = 1000
    for fi47 := 0; fi47 < 14; fi47++ {
        var v48 int32 = 2147483647
        v3 -= (v46 << (((v1 / ((v44) | 1))) & 63))
        _ = fi47
        _ = v48
    }
    var v49 string = ""
    if false {
        var v50 uint64 = 4294967295
        for fi51 := 0; fi51 < 9; fi51++ {
            for fi52 := 0; fi52 < 9; fi52++ {
                v50 += (^(v50))
                switch fi52 {
                case -629:
                    v49 = v49
                    fmt.Println(v5)
                case 1:
                    v3 = (^(v3))
                case -232:
                    fmt.Println(v3, v2)
                }
                _ = len(v4)
                _ = fi52
            }
            var v53 uint64 = v50
            _ = fi51
            _ = v53
        }
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v46
    _ = v49
    _ = v5
    _ = v6
}
