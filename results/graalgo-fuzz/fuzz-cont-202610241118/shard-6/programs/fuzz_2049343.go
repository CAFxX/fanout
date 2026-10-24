package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
type GenS3 struct { F0 uint; F1 int64; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (-2.25 + 3.14159)
    var v3 uint = 1
    v4 := make(map[int]uint32)
    switch 287 {
    case 653:
        _ = len(v1)
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        switch (("\n\t" + ("hello" + "\n\t")) + "") {
        case "\n\t":
            var v6 string = "hello"
            _ = v6
        case "hello":
            fmt.Println(fi5, v2)
        default:
            var v7 uint32 = 131
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 string = "hello"
        _ = fi8
        _ = v9
    }
    var v10 float32 = float32(-2.25)
    v3 = 4294967295
    v11 := GenS1{F0: 2147483648, F1: 0}
    v2 += (float32((0.1 + (3.14159 * 0.0))) - v10)
    fmt.Println(v10)
    if (9223372036854775807 >= -787) {
        var v12 uint8 = 229
        if false {
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := []int{fi13}
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                for fi16 := 0; fi16 < 3; fi16++ {
                    if false {
                        v17 := GenS2{F0: 128, F1: (^(v12)), F2: "hello"}
                        _ = v17
                    }
                    _ = fi16
                }
                v18 := GenS2{F0: v12, F1: v12, F2: (("hello" + (("" + "a") + ("\n\t" + "a"))) + "hello")}
                _ = fi15
                _ = v18
            }
        } else {
            v19 := []int{1}
            var v20 int8 = -1
            var v21 float64 = float64(v2)
            _ = v19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            switch (("\n\t" + (("foo bar" + "hello") + "")) + "a") {
            case "hello":
                v3 = (v3 / ((v3) | 1))
            case "":
                v12 = v12
            }
            v23 := GenS3{F0: (v3 & v3), F1: 1, F2: 429}
            var v24 bool = (264 == 2)
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = v12
    } else {
        v10 -= v10
        delete(v4, -967)
        if true {
            v10 -= v10
            for fi25 := 0; fi25 < 3; fi25++ {
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 uint32 = 65535
                    if true {
                        var v28 int8 = -65
                        var v29 int32 = 885
                        v30 := []int{}
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
                delete(v1, (((("a" + ("\n\t" + "foo bar")) + "\n\t") + ("\n\t" + "hello")) + ("a" + "a")))
                switch 128 {
                case 884:
                    var v31 uint16 = 920
                    _ = v31
                }
                _ = fi25
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                v33 := GenS4{F0: -538}
                v34 := []int{}
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    }
    fmt.Println(v3, v3)
    v35 := GenS1{F0: 801, F1: 201}
    switch -361 {
    case 9223372036854775807:
        v36 := GenS5{F0: false, F1: 42}
        _ = v36
    case 656:
        if (0.0 < -2.25) {
            for fi37 := 0; fi37 < 11; fi37++ {
                for fi38 := 0; fi38 < 20; fi38++ {
                    var v39 uint64 = 484
                    _ = fi38
                    _ = v39
                }
                if true {
                    switch "\n\t" {
                    case "hello":
                        var v40 uint32 = 2
                        _ = v40
                    }
                    var v41 int16 = 1
                    v42 := []int{-9223372036854775808}
                    _ = v41
                    _ = v42
                } else {
                    var v43 bool = true
                    _ = v43
                }
                fmt.Println(v2, fi37)
                _ = fi37
            }
        }
    default:
        v44 := v4[2147483647]
        _ = v44
    }
    if (((1.5 == (0.1 + (1.5 / (-2.25 / (1.5 - (0.1 * 3.14159)))))) || (!true)) && true) {
        v45 := []int{-38, -605, -1}
        _ = v45
    }
    switch "foo bar" {
    case "a":
        v3 = 0
    }
    switch 2147483647 {
    case 32768:
        for fi46 := 0; fi46 < 6; fi46++ {
            fmt.Println(v2, v2, v2)
            var v47 uint64 = 42
            v48 := v1["hello"]
            _ = fi46
            _ = v47
            _ = v48
        }
    default:
        v2 = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
