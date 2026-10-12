package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 uint64, p1 uint16) uint64 {
    if ((0 <= 2) && false) {
        var v4 uint32 = 2
        _ = v4
    } else {
        fmt.Println(p0, p1)
    }
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{(-(fi5)), 748, 1, -2147483648}
        var v7 uint32 = 2147483648
        var v8 int8 = -110
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if (-30 == 129) {
        v9 := GenS1{F0: 293, F1: "a"}
        var v10 uint8 = 44
        for fi11 := 0; fi11 < 16; fi11++ {
            if ((v10 >> ((458) & 63)) != v10) {
                v12 := GenS1{F0: fi11, F1: "\n\t"}
                _ = v12
            }
            p1 = p1
            var v13 int32 = 743
            _ = fi11
            _ = v13
        }
        _ = v10
        _ = v9
    }
    var v14 int8 = 0
    _ = p0
    _ = p1
    _ = v14
    return 4294967295
}
func genFn1(p0 uint16) uint8 {
    if false {
        var v15 uint = uint(p0)
        var v16 int8 = 127
        _ = v15
        _ = v16
    }
    var v17 bool = true
    if v17 {
        v18 := GenS2{F0: false}
        v19 := []int{-2147483648, 845}
        var v20 uint64 = 1
        _ = v18
        _ = v19
        _ = v20
    } else {
        if v17 {
            v21 := []int{-1}
            _ = v21
        }
        if false {
            v22 := GenS3{F0: 65535, F1: ("hello" + "hello")}
            _ = v22
        }
    }
    fmt.Println(v17, v17)
    _ = p0
    _ = v17
    return 127
}

func main() {
    var v1 uint32 = 889
    var v2 uint = 2
    var v3 uint = v2
    v23 := GenS1{F0: 1, F1: ""}
    if (!(genFn1(255) >= 56)) {
        fmt.Println(v1, v2, v2)
        var v24 int16 = 1000
        var v25 uint = v3
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        fmt.Println(v1, fi26)
        if true {
            if true {
                v2 -= 908
                var v27 uint64 = 18446744073709551615
                v28 := GenS4{F0: (fi26 >> ((255) & 63)), F1: fi26, F2: (v1 >> ((v2) & 63))}
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 int8 = 127
                v31 := GenS2{F0: true}
                _ = fi29
                _ = v30
                _ = v31
            }
            if false {
                switch ("a" + "hello") {
                case "a":
                    v1 += (^(v1))
                }
                fmt.Println(v3, v2, v2)
            } else {
                for fi32 := 0; fi32 < 14; fi32++ {
                    v2 -= v2
                    var v33 uint16 = 255
                    switch "\n\t" {
                    case "":
                        v34 := []int{(fi32 >> (((v2 % (((^(v3))) | 1))) & 63)), fi32}
                        _ = v34
                    case "foo bar":
                        v2 += 18446744073709551615
                    }
                    _ = fi32
                    _ = v33
                }
            }
        }
        _ = fi26
    }
    if false {
        var v35 uint16 = 876
        v35 -= (v35 * 1)
        _ = v35
    }
    if true {
        fmt.Println(v1, v3)
    }
    switch "" {
    case "a":
        if (!true) {
            var v36 uint32 = (^(v1))
            v37 := []int{1, -9223372036854775808, 9223372036854775807, 4294967295}
            _ = v36
            _ = v37
        } else {
            v3 += (v3 + v3)
            for fi38 := 0; fi38 < 14; fi38++ {
                v2 = v3
                fmt.Println(v2, fi38)
                fmt.Println(v1)
                _ = fi38
            }
            if ((!true) || (!((65535 >= 1) || (-1 < -32768)))) {
                var v39 float32 = -2.25
                fmt.Println(v2, v2, v2)
                _ = v39
            }
        }
    default:
        v1 = v1
    }
    var v40 int64 = 0
    var v41 int16 = -61
    var v42 uint8 = 2
    switch ("\n\t" + "hello") {
    case "foo bar":
        v43 := []int{-2147483648, -835, 32767, int(v42)}
        _ = v43
    }
    fmt.Println(v42)
    v44 := []int{0}
    var v45 string = "\n\t"
    v45 = ""
    v46 := GenS2{F0: true}
    var v47 float32 = 0.1
    v3 += (v3 * 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
