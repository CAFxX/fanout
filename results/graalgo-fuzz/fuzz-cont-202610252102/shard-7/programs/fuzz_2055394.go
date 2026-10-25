package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64, p1 int16, p2 bool) uint16 {
    p1 -= -32768
    switch 0 {
    case -2147483648:
        var v4 uint64 = p0
        _ = v4
    case 42:
        var v5 int32 = -893
        _ = v5
    default:
        v6 := []int{}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1(p0 uint64, p1 int32, p2 bool) int8 {
    v7 := GenS1{F0: 0}
    fmt.Println(p0, p0, p0)
    var v8 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return int8(v8)
}

func main() {
    var v1 uint = 842
    var v2 int8 = 0
    var v3 int32 = 42
    v1 = v1
    if (true && (!((true && (127 == 128)) && (-258 >= 9223372036854775807)))) {
        var v9 int8 = int8(v1)
        _ = v9
    } else {
        if (v2 == (v2 % ((v2) | 1))) {
            v10 := []int{128, -213, 2147483647, -9223372036854775808}
            _ = v10
        }
        var v11 uint16 = uint16(v1)
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 float64 = 3.14159
            if false {
                var v15 uint16 = 0
                var v16 uint16 = 0
                _ = v15
                _ = v16
            } else {
                var v17 string = "hello"
                var v18 uint64 = 18446744073709551615
                fmt.Println(v18)
                _ = v17
                _ = v18
            }
            var v19 uint64 = 1
            _ = fi13
            _ = v14
            _ = v19
        }
        switch (fi12 >> ((1) & 63)) {
        case -77:
            v2 = -30
        case 1:
            fmt.Println(fi12, fi12, v2)
        default:
            if false {
                var v20 uint16 = 32767
                for fi21 := 0; fi21 < 1; fi21++ {
                    v22 := GenS1{F0: 1}
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
        }
        _ = fi12
    }
    v3 = (v3 * v3)
    var v23 float64 = (1.5 - float64(3.14159))
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(v23, fi24)
        _ = fi24
    }
    v25 := []int{int(v2), -9223372036854775808, -511, 284}
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS1{F0: 2147483647}
        v23 += v23
        for fi28 := 0; fi28 < 4; fi28++ {
            if false {
                switch "" {
                case "":
                    v3 = v3
                    v1 += v1
                case "a":
                    v1 += v1
                case "hello":
                    v25 = append(v25, fi28)
                default:
                    if false {
                        switch (fi28 + 4294967295) {
                        case 2147483647:
                            if (-9223372036854775808 > -281) {
                                v23 = 1.5
                            }
                        }
                        switch "a" {
                        case "":
                            if true {
                                v1 = (v1 << ((v1) & 63))
                            } else {
                                v3 -= int32(v1)
                                _ = len(v25)
                                v29 := GenS1{F0: 241}
                                _ = v29
                            }
                        }
                    }
                }
            }
            v2 += (v2 / (((v2 * v2)) | 1))
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        if len(v25) > 0 {
            v25[0] = fi30
        }
        v3 = v3
        var v31 int64 = 0
        _ = fi30
        _ = v31
    }
    fmt.Println(v23, v2, v3)
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 uint = (^(v1))
        v34 := GenS2{F0: (v3 << (((^(v33))) & 63)), F1: v1, F2: ""}
        fmt.Println(v2)
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            if false {
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                var v37 uint64 = uint64(v1)
                v2 = -1
                _ = v37
            } else {
                var v38 uint8 = 255
                v39 := GenS3{F0: (-(v3))}
                var v40 uint = (^(v1))
                _ = v38
                _ = v39
                _ = v40
            }
            var v41 float64 = v23
            _ = fi36
            _ = v41
        }
        fmt.Println(v2)
        v1 = v1
        _ = fi35
    }
    var v42 int16 = -322
    if false {
        fmt.Println(v1)
        v42 += v42
    } else {
        v23 = v23
        v42 = 1000
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    }
    var v43 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v42
    _ = v43
}
