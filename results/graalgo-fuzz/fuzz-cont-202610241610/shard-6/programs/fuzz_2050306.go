package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    v5 := GenS1{F0: 914}
    var v6 int16 = 67
    var v7 uint32 = 42
    for fi8 := 0; fi8 < 8; fi8++ {
        switch fi8 {
        case -378:
            v9 := []int{127, fi8, fi8}
            _ = v9
        default:
            v10 := []int{fi8, (-(fi8)), fi8}
            _ = v10
        }
        _ = fi8
    }
    if false {
        fmt.Println(v6, v6)
        v11 := []int{26, -370, -355}
        _ = v11
    }
    _ = v5
    _ = v6
    _ = v7
    return (v7 | 2147483648)
}

func main() {
    var v1 uint32 = 838
    var v2 uint64 = 812
    var v3 uint64 = (v2 % (((v2 - v2)) | 1))
    var v4 int8 = -1
    var v12 int8 = v4
    switch ("a" + ("\n\t" + ((("" + ("\n\t" + ("hello" + ("\n\t" + ("\n\t" + "a"))))) + "foo bar") + "\n\t"))) {
    case "a":
        var v13 bool = false
        _ = v13
    }
    v12 -= (^(v12))
    if false {
        var v14 uint = 668
        v3 -= (v2 | v2)
        _ = v14
    } else {
        switch 1 {
        case -65:
            fmt.Println(v1, v2)
        case 4294967295:
            v15 := GenS2{F0: (^(v1))}
            _ = v15
        case 32767:
            v4 += (v4 & 100)
        default:
            var v16 float32 = float32(1.5)
            _ = v16
        }
        var v17 string = "hello"
        _ = v17
    }
    v18 := GenS2{F0: (v1 >> ((1) & 63))}
    if (-9223372036854775808 > 42) {
        v19 := []int{128, 32767, 32767, 32767}
        var v20 uint64 = v3
        var v21 uint8 = 1
        _ = v19
        _ = v20
        _ = v21
    } else {
        var v22 float64 = (1.5 * 1e10)
        _ = v22
    }
    fmt.Println(v3, v3)
    v23 := GenS3{F0: 0}
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 int8 = v4
        v26 := []int{(fi24 << ((904) & 63)), (-(fi24)), fi24, (fi24 >> ((42) & 63))}
        if true {
            v27 := GenS2{F0: v1}
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 int32 = 311
                var v30 uint64 = (^(v3))
                _ = fi28
                _ = v29
                _ = v30
            }
            var v31 int64 = -1
            _ = v27
            _ = v31
        } else {
            if true {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v33 := GenS3{F0: 127}
                    v25 = (v12 % (((-(v12))) | 1))
                    v4 -= 0
                    _ = fi32
                    _ = v33
                }
                v1 = (v1 & (^(v1)))
                if (v3 != 42) {
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        var v34 float64 = -2.25
                        _ = v34
                    case "a":
                        v26 = append(v26, -324)
                    }
                } else {
                    var v35 uint32 = (v1 - (^(v1)))
                    if false {
                        v1 += v35
                        v1 += 2
                        v35 += uint32(v3)
                    }
                    _ = v35
                }
            }
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println(v2, v1, v4)
    var v36 string = "hello"
    var v37 uint8 = 177
    var v38 uint32 = (v1 + 2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
