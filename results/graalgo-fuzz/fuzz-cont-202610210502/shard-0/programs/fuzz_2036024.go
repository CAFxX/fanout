package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint; F2 string }

func main() {
    var v1 string = ("" + ("\n\t" + "foo bar"))
    var v2 uint = 598
    var v3 int64 = 42
    var v4 int = -2147483648
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{(fi5 + (-(v4))), 9223372036854775807}
        fmt.Println(v1)
        var v7 int8 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    switch v4 {
    case 0:
        if false {
            if true {
                v2 -= v2
                fmt.Println(v3, v4, v2)
                for fi8 := 0; fi8 < 9; fi8++ {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        if true {
                            v3 -= v3
                        }
                        _ = fi9
                    }
                    switch 127 {
                    case 32767:
                        fmt.Println(v2, v2)
                    case -2147483648:
                        var v10 string = ""
                        _ = v10
                    }
                    _ = fi8
                }
            } else {
                v11 := []int{v4, (v4 % ((v4) | 1)), (v4 ^ (^(v4)))}
                _ = v11
            }
            if (!true) {
                var v12 uint16 = 42
                switch (v4 >> (((v2 + (v2 >> ((v2) & 63)))) & 63)) {
                case 0:
                    v13 := GenS1{F0: v4}
                    _ = v13
                case 9223372036854775807:
                    var v14 int16 = -32768
                    _ = v14
                }
                var v15 int = v4
                _ = v12
                _ = v15
            } else {
                if (v1 <= v1) {
                    var v16 int32 = 42
                    switch (^(v4)) {
                    case 32767:
                        v17 := []int{876}
                        _ = v17
                    case 152:
                        for fi18 := 0; fi18 < 16; fi18++ {
                            var v19 uint64 = 0
                            _ = fi18
                            _ = v19
                        }
                    }
                    v20 := GenS1{F0: (v4 / ((v4) | 1))}
                    _ = v16
                    _ = v20
                }
            }
        }
    }
    switch v4 {
    case 32768:
        var v21 string = "hello"
        _ = v21
    default:
        v4 = v4
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 uint = 454
        if (!(((2 < (v23 << ((v23) & 63))) && (v23 < v23)) || (!((!false) || (v4 >= fi22))))) {
            switch v1 {
            case "hello":
                v24 := GenS2{F0: 4294967295}
                _ = v24
            }
            v3 += v3
        } else {
            v25 := GenS3{F0: 1, F1: 329, F2: v1}
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 int32 = 42
                _ = fi26
                _ = v27
            }
            var v28 int8 = 100
            _ = v25
            _ = v28
        }
        _ = fi22
        _ = v23
    }
    var v29 uint64 = 168
    for fi30 := 0; fi30 < 12; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            v32 := []int{v4, fi30, 919}
            var v33 uint32 = uint32(v2)
            _ = fi31
            _ = v32
            _ = v33
        }
        switch v1 {
        case "hello":
            if (!(!(((-9223372036854775808 != -1) && (true || (true || false))) || false))) {
                if (50 >= 255) {
                    var v34 string = v1
                    _ = v34
                } else {
                    v35 := []int{v4, v4, 32767, fi30}
                    v4 = 4294967295
                    v2 = 18446744073709551615
                    _ = v35
                }
                v29 = (v29 | v29)
            }
        default:
            v29 += (v29 & v29)
        }
        _ = fi30
    }
    fmt.Println(v3, v1)
    var v36 uint32 = 0
    v37 := []int{v4}
    fmt.Println(v2)
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 int = (v4 + (v4 | 32768))
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
