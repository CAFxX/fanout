package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint16 = 32767
    var v2 bool = (false && (1 < -128))
    var v3 uint32 = 254
    var v4 bool = false
    for fi5 := 0; fi5 < 3; fi5++ {
        v4 = (int16(v1) <= -32768)
        fmt.Println(v1)
        switch (-(fi5)) {
        case -9223372036854775808:
            var v6 bool = true
            _ = v6
        case -100:
            var v7 int64 = 599
            _ = v7
        case 32768:
            v8 := GenS1{F0: 42, F1: v3}
            _ = v8
        default:
            v1 = 157
        }
        _ = fi5
    }
    v2 = (v2 || (0 >= 0))
    for fi9 := 0; fi9 < 16; fi9++ {
        if v4 {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 string = ""
                v2 = (!(42 < 127))
                _ = fi10
                _ = v11
            }
            fmt.Println(v4, v2)
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            if (!false) {
                if false {
                    var v13 int16 = -32768
                    var v14 float32 = -2.25
                    var v15 float32 = v14
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v1 += (^(v1))
                v16 := []int{(fi12 + (fi9 * (fi9 / ((fi12) | 1)))), -9223372036854775808, (fi9 + fi12)}
                _ = v16
            } else {
                fmt.Println(v1, v3)
                var v17 int32 = 2147483647
                if true {
                    v18 := GenS1{F0: -896, F1: v3}
                    _ = v18
                } else {
                    if v2 {
                        v2 = ((-160 >= fi9) || v2)
                    } else {
                        v19 := GenS1{F0: 9223372036854775807, F1: (v3 / (((v3 >> ((uint(v1)) & 63))) | 1))}
                        _ = v19
                    }
                    v3 = (v3 ^ v3)
                    var v20 string = ("foo bar" + "\n\t")
                    var v21 float32 = 1.5
                    fmt.Println(v1)
                    var v22 int = fi9
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = v17
            }
            if (808 > (v1 ^ (v1 + v1))) {
                var v23 int16 = 458
                v1 = v1
                v2 = (true || true)
                v1 += uint16(v3)
                v4 = v4
                _ = v23
            }
            _ = fi12
        }
        _ = fi9
    }
    v4 = false
    var v24 int = 32767
    v25 := []int{(v24 + (v24 - v24))}
    fmt.Println(v2, v1)
    if v4 {
        switch "hello" {
        case "hello":
            v24 -= 0
        case "":
            var v26 uint16 = v1
            _ = v26
        default:
            _ = len(v25)
        }
        switch (("\n\t" + ("hello" + "")) + ("a" + "")) {
        case "hello":
            var v27 uint64 = 165
            _ = v27
        }
    }
    v25 = append(v25, v24)
    if v2 {
        if v2 {
            var v28 uint64 = 18446744073709551615
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = v28
        }
        var v29 uint8 = 128
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            if v4 {
                if (-32768 >= -1) {
                    if (!v2) {
                        v2 = (42 == 1)
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        _ = len(v25)
                    } else {
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        v1 = v1
                    }
                } else {
                    var v32 int = (^(fi31))
                    v25 = append(v25, -9223372036854775808)
                    fmt.Println(v1, fi31, fi30)
                    _ = v32
                }
                var v33 bool = true
                _ = v33
            } else {
                var v34 int32 = -2147483648
                switch (^(fi31)) {
                case 9223372036854775807:
                    var v35 float32 = 0.1
                    _ = v35
                }
                v1 = 255
                _ = v34
            }
            _ = fi31
        }
        var v36 float64 = float64(1.5)
        _ = fi30
        _ = v36
    }
    v4 = ((^(v3)) > 4294967295)
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 int16 = 42
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    v1 += (v1 - (v1 % (((v1 & v1)) | 1)))
    var v40 bool = ((!(!(!(v4 && (1e10 <= 0.1))))) && ((!(-1 >= -1)) || true))
    if (true || ((1 <= 1) && true)) {
        var v41 bool = false
        v1 = 2
        _ = v41
    }
    _ = len(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v40
}
