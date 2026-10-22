package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int16, p1 string) bool {
    var v4 int8 = 127
    if true {
        switch 1 {
        case 9223372036854775807:
            var v5 string = p1
            _ = v5
        }
    } else {
        p0 = 1
        v6 := GenS1{F0: 1}
        if true {
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 uint = 631
                for fi9 := 0; fi9 < 10; fi9++ {
                    v10 := []int{fi7, int(p0), fi9}
                    fmt.Println(v8, fi7, p0)
                    _ = fi9
                    _ = v10
                }
                switch 9223372036854775807 {
                case 42:
                    v11 := GenS1{F0: 369}
                    _ = v11
                case -1:
                    var v12 int16 = 1
                    _ = v12
                case -2147483648:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v14 := []int{-2147483648}
                        _ = fi13
                        _ = v14
                    }
                default:
                    var v15 uint32 = 65535
                    _ = v15
                }
                _ = fi7
                _ = v8
            }
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return (1 == 42)
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint64 = 0
    v3 := make(map[string]string)
    delete(v3, "foo bar")
    if false {
        v3["a"] = (("" + ("hello" + "")) + "")
        fmt.Println(v1, v1)
        switch "hello" {
        case "":
            if (!true) {
                v16 := GenS1{F0: uint64(v1)}
                _ = v16
            } else {
                var v17 uint8 = 42
                var v18 int64 = 0
                v1 = 65535
                _ = v17
                _ = v18
            }
        }
    } else {
        var v19 uint8 = 42
        _ = v19
    }
    _ = len(v3)
    v2 = 42
    v1 += (v1 & v1)
    v20 := GenS1{F0: (v2 / ((441) | 1))}
    v2 = v2
    var v21 uint32 = (^(v1))
    for fi22 := 0; fi22 < 16; fi22++ {
        if (0 >= int64(fi22)) {
            v23 := v3["foo bar"]
            _ = v23
        }
        _ = fi22
    }
    var v24 uint = 255
    for fi25 := 0; fi25 < 20; fi25++ {
        v2 += (v2 + 125)
        v24 = v24
        if (v2 <= 42) {
            switch (fi25 << ((2) & 63)) {
            case 128:
                for fi26 := 0; fi26 < 5; fi26++ {
                    v27 := []int{(fi25 & (fi25 * fi26)), (fi26 % (((fi26 % ((9223372036854775807) | 1))) | 1))}
                    v28 := GenS2{F0: 65}
                    switch fi25 {
                    case 4294967295:
                        v29 := GenS1{F0: 42}
                        _ = v29
                    case 911:
                        switch (^(fi25)) {
                        case 0:
                            if (!true) {
                                v27 = append(v27, (fi26 | fi25))
                                var v30 int64 = 0
                                _ = v30
                            }
                        }
                    }
                    fmt.Println(v24, v1)
                    var v31 float32 = 0.1
                    fmt.Println(v21)
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v31
                }
            case 2147483647:
                v21 = 65535
            default:
                var v32 bool = (uint64(v21) != 661)
                _ = v32
            }
            var v33 float32 = (0.0 - 3.14159)
            switch (fi25 - 782) {
            case 4294967295:
                delete(v3, (("hello" + "foo bar") + "\n\t"))
            case 128:
                var v34 uint64 = (v2 << ((v24) & 63))
                _ = v34
            default:
                v35 := []int{}
                _ = v35
            }
            _ = v33
        }
        _ = fi25
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        for fi37 := 0; fi37 < 12; fi37++ {
            v21 -= v1
            var v38 int16 = 32767
            _ = fi37
            _ = v38
        }
        var v39 int = fi36
        if (false && true) {
            var v40 uint64 = 2
            _ = v40
        }
        _ = fi36
        _ = v39
    }
    var v41 float64 = 0.0
    var v42 float32 = 3.14159
    switch 1 {
    case 244:
        fmt.Println(v42, v2)
    default:
        v43 := GenS3{F0: v2, F1: (v2 + (v2 - v2)), F2: 1}
        _ = v43
    }
    v44 := GenS4{F0: uint8(v21)}
    v45 := []int{-9223372036854775808, 127, 869}
    var v46 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v46
}
