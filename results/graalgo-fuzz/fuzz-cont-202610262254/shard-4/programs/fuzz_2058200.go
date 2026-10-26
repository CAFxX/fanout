package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0(p0 uint16, p1 uint) int8 {
    switch (("" + "a") + ("\n\t" + "hello")) {
    case "foo bar":
        if false {
            fmt.Println(p0, p0)
            p0 = 42
        }
    default:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int64 = 1
            if ("" != "a") {
                v9 := GenS1{F0: (v8 / ((int64(fi7)) | 1))}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    v10 := GenS1{F0: 9223372036854775807}
    if (127 < -70) {
        v11 := GenS2{F0: 1}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(fi12)
                v14 := []int{(fi13 | 4294967295), fi13, -9223372036854775808}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        v15 := []int{-1, 0, 456}
        p0 = (p0 / ((p0) | 1))
        _ = v15
    }
    if ((((1 > 0) || (p0 <= (p0 & 255))) && (914 <= 332)) || true) {
        v16 := []int{32768, -89, 42}
        var v17 uint = p1
        _ = v16
        _ = v17
    } else {
        p1 += p1
        var v18 float64 = 3.14159
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    v5 := []string{v4, v4, v4}
    var v6 int = 2147483647
    var v19 int16 = -1
    var v20 int8 = 127
    var v21 int16 = v19
    v21 -= 0
    for fi22 := 0; fi22 < 14; fi22++ {
        if true {
            for fi23 := 0; fi23 < 5; fi23++ {
                v3[65535] = (-1 <= 9223372036854775807)
                _ = len(v5)
                _ = fi23
            }
            v24 := GenS2{F0: v1}
            _ = v24
        }
        switch v4 {
        case "foo bar":
            var v25 int8 = v20
            _ = v25
        case "a":
            var v26 uint8 = 2
            _ = v26
        case "":
            switch 104 {
            case 32767:
                var v27 uint16 = (v1 % (((v1 - (v1 / ((0) | 1)))) | 1))
                _ = v27
            case -1:
                v1 = v1
            case 2147483647:
                v1 -= 65535
            }
        }
        switch int(v1) {
        case 589:
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := GenS2{F0: v1}
                _ = fi28
                _ = v29
            }
        case 32768:
            if (false && false) {
                switch (v4 + ("a" + v4)) {
                case "":
                    v19 += v21
                case "hello":
                    var v30 bool = (0 == -2147483648)
                    _ = v30
                }
                var v31 int32 = 0
                _ = v31
            }
        default:
            fmt.Println(v21, v20)
        }
        _ = fi22
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 int = fi32
        _ = fi32
        _ = v33
    }
    var v34 bool = false
    if len(v5) > 0 {
        v5[0] = (v4 + v4)
    }
    _ = len(v2)
    v20 += genFn0(1, 2)
    if (2 != v1) {
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := GenS1{F0: 0}
            v4 = ("hello" + v4)
            for fi37 := 0; fi37 < 3; fi37++ {
                v38 := GenS3{F0: 42, F1: v34, F2: 995}
                fmt.Println(v21, fi35)
                _ = fi37
                _ = v38
            }
            _ = fi35
            _ = v36
        }
    }
    var v39 int = -1
    var v40 uint64 = 0
    fmt.Println(v1, v39)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
