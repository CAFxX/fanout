package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint, p2 int32) string {
    p2 = p2
    p2 = 0
    v4 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ((("a" + "hello") + "\n\t") + "a")
}
func genFn1(p0 string, p1 uint16) int64 {
    if false {
        for fi5 := 0; fi5 < 16; fi5++ {
            fmt.Println(p1, fi5, p1)
            v6 := []int{-209, (-(fi5)), (fi5 + fi5)}
            fmt.Println(fi5)
            _ = fi5
            _ = v6
        }
        p1 = (p1 - p1)
    } else {
        var v7 bool = true
        _ = v7
    }
    var v8 float64 = 0.1
    var v9 string = ("a" + "\n\t")
    fmt.Println(p0)
    switch (v9 + ((((v9 + "hello") + (((((((v9 + ((v9 + (p0 + p0)) + v9)) + "foo bar") + v9) + "hello") + v9) + "foo bar") + ("\n\t" + "foo bar"))) + ((v9 + ((v9 + (p0 + p0)) + (((v9 + "") + (v9 + p0)) + "foo bar"))) + "")) + (v9 + "a"))) {
    case "hello":
        if false {
            if false {
                var v10 int8 = 0
                _ = v10
            }
        }
    case "":
        var v11 int = 32767
        _ = v11
    default:
        v9 = "foo bar"
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint16 = 2
    var v2 float32 = (float32(-2.25) + 3.14159)
    var v3 bool = false
    if (2147483647 != 4294967295) {
        var v12 int16 = -32768
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint = 42
            v12 -= -1
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    var v15 uint32 = 290
    var v16 uint32 = v15
    var v17 int16 = 32767
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint8 = 119
        var v20 uint16 = (v1 & (^(v1)))
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{481, -2147483648, int(v1)}
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint8 = 207
        switch "foo bar" {
        case "hello":
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint8 = v23
                v25 = v25
                _ = fi24
                _ = v25
            }
        case "":
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        case "foo bar":
            switch fi22 {
            case -662:
                fmt.Println(v3)
            }
        }
        _ = fi22
        _ = v23
    }
    if (true || true) {
        var v26 uint16 = 65535
        v27 := GenS2{F0: 610, F1: 0}
        var v28 uint16 = v1
        _ = v26
        _ = v27
        _ = v28
    } else {
        switch "\n\t" {
        case "\n\t":
            if len(v21) > 0 {
                v21[0] = 2147483647
            }
        }
    }
    fmt.Println(v2, v1, v15)
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint64 = 2
        _ = fi29
        _ = v30
    }
    var v31 string = "a"
    v17 = (v17 << ((2) & 63))
    v15 += 2147483647
    if v3 {
        v17 -= int16(v1)
        switch v31 {
        case "foo bar":
            v32 := GenS3{F0: (-(v17))}
            _ = v32
        case "hello":
            var v33 uint8 = 129
            _ = v33
        case "a":
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        default:
            v21 = append(v21, 952)
        }
    }
    var v34 int64 = 1
    switch v31 {
    case "":
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 float32 = v2
            var v37 int8 = -128
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 int8 = v37
                v21 = append(v21, (fi38 ^ (fi35 - int(v34))))
                fmt.Println(v3, v16, fi38)
                _ = fi38
                _ = v39
            }
            _ = fi35
            _ = v36
            _ = v37
        }
    case "foo bar":
        switch -1 {
        case 32767:
            if len(v21) > 0 {
                v21[0] = 989
            }
        case -2147483648:
            v3 = (true && (v3 && (false || ((v34 | v34) > genFn1("", 1)))))
        }
    }
    var v40 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v34
    _ = v40
}
