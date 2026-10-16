package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 65535
    for fi10 := 0; fi10 < 9; fi10++ {
        if (-128 != -122) {
            for fi11 := 0; fi11 < 17; fi11++ {
                if false {
                    v9 = v9
                    var v12 float32 = float32(((float64(-2.25) + 1e10) * (0.1 - 0.0)))
                    _ = v12
                }
                _ = fi11
            }
            v9 += 0
            var v13 bool = (false || false)
            _ = v13
        } else {
            v14 := []int{(fi10 << ((42) & 63))}
            _ = v14
        }
        switch fi10 {
        case -719:
            var v15 string = ("foo bar" + ("hello" + "\n\t"))
            _ = v15
        case 42:
            switch ("" + "") {
            case "":
                var v16 uint16 = 0
                _ = v16
            }
        }
        _ = fi10
    }
    if (32767 < 9223372036854775807) {
        switch -1 {
        case 128:
            var v17 int64 = 0
            _ = v17
        case -971:
            var v18 uint64 = 4294967295
            _ = v18
        case 1:
            v19 := []int{-862, 32767, -1}
            _ = v19
        default:
            v20 := GenS1{F0: 255, F1: 100, F2: 782}
            _ = v20
        }
        var v21 uint8 = 42
        _ = v21
    }
    var v22 float32 = 0.1
    _ = v22
    _ = v9
    return 73
}
func genFn1() int {
    v23 := GenS2{F0: -565, F1: (!false)}
    v24 := []int{-176, 32768, -2147483648, 439}
    _ = v23
    _ = v24
    return -2147483648
}

func main() {
    var v1 int8 = 100
    var v2 uint16 = 1
    var v3 float64 = (0.0 / 1.0)
    v4 := make(map[int]bool)
    var v5 int8 = (v1 / (((v1 + (v1 / ((0) | 1)))) | 1))
    var v6 bool = (-1 > 0)
    var v25 uint8 = 6
    var v26 string = "hello"
    v1 = (v5 % (((v5 - v5)) | 1))
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 19; fi28++ {
            v25 += 41
            _ = fi28
        }
        if (1000 == 1000) {
            var v29 bool = v6
            var v30 bool = v6
            _ = v29
            _ = v30
        } else {
            switch (fi27 >> ((42) & 63)) {
            case 4294967295:
                if v6 {
                    v31 := GenS1{F0: (^(v2)), F1: 42, F2: 2147483647}
                    v32 := GenS1{F0: v2, F1: genFn0(), F2: 17}
                    _ = v31
                    _ = v32
                }
            case 9223372036854775807:
                v26 = (v26 + "\n\t")
            default:
                for fi33 := 0; fi33 < 19; fi33++ {
                    v26 = (v26 + (v26 + (("\n\t" + v26) + ("foo bar" + ("hello" + v26)))))
                    switch (-(fi33)) {
                    case 1:
                        var v34 bool = false
                        _ = v34
                    case 224:
                        var v35 int64 = 42
                        _ = v35
                    case 32768:
                        if v6 {
                            var v36 uint8 = 110
                            _ = v36
                        } else {
                            var v37 bool = v6
                            _ = v37
                        }
                    default:
                        var v38 float64 = v3
                        _ = v38
                    }
                    _ = fi33
                }
            }
            var v39 int64 = 42
            fmt.Println(fi27, v3)
            _ = v39
        }
        _ = fi27
    }
    fmt.Println(v1, v6, v2)
    fmt.Println(v1, v1, v3)
    v5 = genFn0()
    _ = len(v4)
    switch "\n\t" {
    case "\n\t":
        switch -2147483648 {
        case -179:
            v1 += (-(v1))
        case -2147483648:
            var v40 uint16 = 65535
            _ = v40
        }
    }
    v41 := GenS3{F0: int64(v1), F1: 2}
    fmt.Println(v2, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
