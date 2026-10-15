package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 string; F1 bool; F2 int32 }
type GenS5 struct { F0 bool }
func genFn0(p0 int8, p1 bool, p2 uint) uint8 {
    var v6 uint64 = 0
    v7 := GenS1{F0: 2}
    switch ("a" + ("foo bar" + "foo bar")) {
    case "\n\t":
        switch "a" {
        case "foo bar":
            var v8 int8 = 100
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 bool = (p1 || (!(357 < -32768)))
                _ = fi9
                _ = v10
            }
        case "":
            p2 += (p2 % (((p2 & p2)) | 1))
        default:
            var v11 bool = p1
            _ = v11
        }
    case "foo bar":
        if ("" == (((("hello" + ("foo bar" + "a")) + ("hello" + ((("a" + "\n\t") + ("a" + ("\n\t" + ""))) + ("" + (("\n\t" + (("a" + "") + ("" + "a"))) + "foo bar"))))) + ("hello" + "hello")) + ("" + ("hello" + "hello")))) {
            var v12 bool = ((false && (1.5 >= 1e10)) || false)
            switch "\n\t" {
            case "\n\t":
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := []int{32767}
                    _ = fi13
                    _ = v14
                }
            default:
                var v15 int32 = -2147483648
                _ = v15
            }
            _ = v12
        }
    case "hello":
        fmt.Println(p1, p0, p1)
    }
    fmt.Println(p1, v6, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 uint8, p2 uint8) int {
    var v16 int32 = 42
    switch ("" + ("hello" + "hello")) {
    case "a":
        if ((p2 & genFn0(-128, (true || (1 < 91)), (p0 - p0))) > (p1 % ((p1) | 1))) {
            p1 = (p1 << ((uint(p1)) & 63))
            v17 := []int{}
            _ = v17
        }
    case "":
        p2 = p2
    case "foo bar":
        var v18 float32 = float32(-2.25)
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = "foo bar"
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 float64 = 0.1
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int64 = 1
        var v21 float32 = (-2.25 * 1.5)
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := []int{(fi22 / (((fi22 | (^(fi19)))) | 1))}
            _ = fi22
            _ = v23
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v24 := GenS1{F0: genFn0(0, (1 < 100), 0)}
    fmt.Println(v5)
    v1 = 2
    var v25 uint8 = genFn0(42, (!(!false)), 0)
    if (!true) {
        _ = len(v4)
        var v26 uint16 = 65535
        var v27 uint32 = v1
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS2{F0: 42, F1: (-(fi28))}
        v25 = v25
        _ = fi28
        _ = v29
    }
    if false {
        if ((^(v1)) < (^(v1))) {
            v30 := GenS3{F0: v1, F1: 0}
            _ = v30
        } else {
            var v31 uint32 = 732
            v32 := GenS1{F0: v25}
            v33 := GenS4{F0: v2, F1: (0 >= 42), F2: 0}
            _ = v31
            _ = v32
            _ = v33
        }
    } else {
        var v34 int = 123
        fmt.Println(v1, v5, v25)
        fmt.Println(v5)
        _ = v34
    }
    v2 = v2
    if ((false || true) || true) {
        var v35 uint = 65535
        var v36 int = genFn1(464, 35, (v25 * 1))
        for fi37 := 0; fi37 < 18; fi37++ {
            fmt.Println(v36)
            _ = fi37
        }
        _ = v35
        _ = v36
    } else {
        if ("foo bar" < v2) {
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 int = fi38
                _ = fi38
                _ = v39
            }
        }
        for fi40 := 0; fi40 < 12; fi40++ {
            fmt.Println(fi40, v1, v1)
            v41 := GenS3{F0: v1, F1: 0}
            v42 := GenS5{F0: (true || false)}
            _ = fi40
            _ = v41
            _ = v42
        }
        var v43 int8 = 127
        _ = v43
    }
    v5 = (((v5 / float64(0.1)) * float64(-2.25)) * ((v5 * float64(float32(((float64(((0.0 / 0.1) * 3.14159)) + (v5 - 1.5)) + (v5 / float64(3.14159)))))) - float64(0.0)))
    if false {
        v44 := []int{}
        var v45 int32 = 42
        for fi46 := 0; fi46 < 12; fi46++ {
            v47 := v4[(v1 ^ v1)]
            var v48 uint8 = 144
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = v44
        _ = v45
    } else {
        var v49 uint = 1
        var v50 uint32 = (v1 >> ((v49) & 63))
        for fi51 := 0; fi51 < 5; fi51++ {
            for fi52 := 0; fi52 < 7; fi52++ {
                v49 = v49
                v53 := v4[v1]
                v53 = "hello"
                _ = fi52
                _ = v53
            }
            _ = fi51
        }
        _ = v49
        _ = v50
    }
    var v54 float64 = (v5 / float64((1.5 - ((float32(v5) + 1.5) / 1.0))))
    var v55 uint8 = (v25 - 127)
    var v56 int = -1
    fmt.Println(v5, v2)
    if true {
        v57 := GenS3{F0: v1, F1: 1}
        v55 = (v55 & v25)
        _ = v57
    }
    switch (v2 + (v2 + v2)) {
    case "\n\t":
        delete(v4, 65535)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
}
