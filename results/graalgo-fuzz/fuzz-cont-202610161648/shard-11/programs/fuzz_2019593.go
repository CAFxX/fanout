package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 string; F2 uint16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int; F1 bool }
func genFn0() int64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 bool = false
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint8 = 127
            switch "a" {
            case "\n\t":
                if (!(1.5 >= -2.25)) {
                    var v9 uint8 = (v8 + (v8 + v8))
                    v8 = v9
                    _ = v9
                } else {
                    var v10 int8 = 127
                    _ = v10
                }
            case "foo bar":
                switch fi4 {
                case 42:
                    v11 := []int{(fi7 | -2147483648), 128, (fi4 % (((fi7 >> ((255) & 63))) | 1))}
                    _ = v11
                case 613:
                    switch (fi4 / (((fi7 % ((fi7) | 1))) | 1)) {
                    case 32767:
                        fmt.Println(fi4, fi7, v8)
                    }
                case 127:
                    var v12 uint = 687
                    _ = v12
                }
            default:
                var v13 int32 = 1
                _ = v13
            }
            _ = fi7
            _ = v8
        }
        _ = fi4
    }
    if (false || true) {
        var v14 string = "foo bar"
        fmt.Println(v14)
        var v15 uint64 = 42
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: -596}
    v17 := []int{895, 2147483647}
    if true {
        var v18 int64 = 385
        _ = v18
    }
    _ = v16
    _ = v17
    return -856
}

func main() {
    v1 := []int{2147483647, 810}
    var v2 int16 = 0
    v3 := []string{"\n\t"}
    if (667 > 18446744073709551615) {
        switch "\n\t" {
        case "a":
            v2 = v2
        case "foo bar":
            fmt.Println(v2)
        }
        if false {
            var v19 bool = true
            if (-997 <= 42) {
                switch -549 {
                case 32768:
                    v20 := GenS2{F0: 926, F1: "a", F2: 255}
                    _ = v20
                case 42:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                var v21 float32 = (-2.25 * 0.1)
                _ = v21
            }
            _ = v19
        } else {
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 7; fi23++ {
                    var v24 bool = (0.1 < -2.25)
                    v3 = append(v3, "hello")
                    _ = fi23
                    _ = v24
                }
                v25 := GenS1{F0: fi22}
                _ = fi22
                _ = v25
            }
            v3 = append(v3, "hello")
        }
        v26 := GenS1{F0: -9223372036854775808}
        _ = v26
    }
    var v27 int64 = 0
    var v28 float32 = 1.5
    v29 := GenS3{F0: (v2 * 1000), F1: (!((((false && false) || false) || (1e10 >= 3.14159)) && false))}
    var v30 uint8 = 128
    var v31 uint16 = 103
    for fi32 := 0; fi32 < 1; fi32++ {
        v33 := GenS1{F0: fi32}
        if (false && false) {
            if ((!(!((^(v31)) >= 42))) || ("\n\t" > (("" + "\n\t") + "foo bar"))) {
                var v34 uint8 = v30
                switch "a" {
                case "foo bar":
                    var v35 float32 = -2.25
                    _ = v35
                case "hello":
                    v2 -= v2
                case "\n\t":
                    if false {
                        if (0 > -37) {
                            var v36 bool = (true && true)
                            v1 = append(v1, (fi32 | 32767))
                            _ = v36
                        }
                    }
                default:
                    var v37 int8 = 42
                    _ = v37
                }
                if true {
                    var v38 string = "foo bar"
                    var v39 uint64 = 0
                    var v40 int = (-(fi32))
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = v34
            }
            var v41 int32 = -1
            _ = v41
        }
        v30 += v30
        _ = fi32
        _ = v33
    }
    v1 = append(v1, -9223372036854775808)
    var v42 bool = false
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 uint32 = 2
        fmt.Println(v42)
        fmt.Println(v42)
        _ = fi43
        _ = v44
    }
    v45 := GenS3{F0: 32767, F1: (-1 >= v2)}
    v46 := GenS4{F0: 2147483647, F1: (!(!true))}
    v28 += 1e10
    var v47 uint = 42
    v47 = (v47 ^ v47)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v42
    _ = v45
    _ = v46
    _ = v47
}
