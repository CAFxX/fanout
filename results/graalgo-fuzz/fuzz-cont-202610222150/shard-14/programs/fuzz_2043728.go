package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint; F1 int8; F2 uint }
func genFn0(p0 int8) uint32 {
    v5 := []int{-961, 0}
    if false {
        p0 = (p0 + 127)
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 uint32 = 65535
                v5 = append(v5, fi6)
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: 32767, F1: 2}
            _ = fi6
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        if false {
            p0 = 127
            _ = len(v5)
            p0 -= (^(p0))
        } else {
            v11 := GenS2{F0: 1}
            _ = v11
        }
        _ = fi10
    }
    if false {
        v12 := GenS3{F0: 135, F1: 96}
        if false {
            p0 = p0
            switch "\n\t" {
            case "\n\t":
                v13 := GenS3{F0: 4294967295, F1: 191}
                _ = v13
            }
            var v14 float64 = (-2.25 - (1e10 - (0.1 + float64(1e10))))
            _ = v14
        } else {
            switch ("foo bar" + "hello") {
            case "":
                switch (("foo bar" + (("" + "") + "a")) + "foo bar") {
                case "foo bar":
                    var v15 int16 = 32767
                    _ = v15
                case "hello":
                    p0 = 96
                }
            case "a":
                var v16 int8 = (p0 << ((1) & 63))
                _ = v16
            case "hello":
                var v17 uint16 = 42
                _ = v17
            default:
                var v18 float64 = (3.14159 / 1.0)
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(p0)
            _ = fi19
        }
        _ = v12
    }
    p0 = (p0 & p0)
    _ = p0
    _ = v5
    return 2
}
func genFn1() uint64 {
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := GenS3{F0: 4294967295, F1: 2}
        for fi22 := 0; fi22 < 8; fi22++ {
            fmt.Println(fi22, fi20)
            v23 := []int{(-(fi20)), 128, (fi20 - fi20)}
            v23 = append(v23, fi22)
            _ = fi22
            _ = v23
        }
        var v24 uint64 = 0
        _ = fi20
        _ = v21
        _ = v24
    }
    v25 := []int{1, 294, 32767, 9223372036854775807}
    _ = v25
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 uint16 = 0
    var v3 int = -9223372036854775808
    var v4 float32 = (0.0 - ((0.1 + 1e10) + ((((0.1 / 1.5) / 1e10) + 1e10) + 1e10)))
    v26 := GenS4{F0: 4294967295, F1: 130}
    if false {
        fmt.Println(v3, v4, v3)
    }
    fmt.Println(v4)
    v27 := GenS3{F0: genFn1(), F1: 34}
    var v28 int8 = 100
    var v29 bool = (4294967295 == v3)
    for fi30 := 0; fi30 < 12; fi30++ {
        switch (^(v3)) {
        case 441:
            if (42 >= 0) {
                v2 = 2
                fmt.Println(fi30, v2, v3)
            }
        case 498:
            if (42 >= 1) {
                v1 -= float64((v4 + v4))
                for fi31 := 0; fi31 < 1; fi31++ {
                    if true {
                        v32 := []int{fi30, (fi30 * -2147483648), fi30, v3}
                        _ = v32
                    }
                    if (1 != 65535) {
                        fmt.Println(v1, v2)
                        var v33 uint32 = 2147483648
                        _ = v33
                    } else {
                        v34 := []int{fi31, v3, v3, 32768}
                        var v35 uint64 = 18446744073709551615
                        _ = v34
                        _ = v35
                    }
                    v28 = v28
                    _ = fi31
                }
            } else {
                v28 = -69
                v28 -= v28
                if v29 {
                    var v36 uint32 = 2147483648
                    _ = v36
                }
            }
        }
        _ = fi30
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 11; fi38++ {
            switch (fi38 + (v3 >> ((0) & 63))) {
            case 32768:
                var v39 float32 = (v4 * (0.0 - (float32(((((1.5 - float64((v4 * v4))) * -2.25) / (float64(v4) * ((v1 - (v1 - (float64(v4) / 1e10))) - v1))) + float64(v4))) * (3.14159 - (0.0 / (((v4 + v4) / -2.25) * (-2.25 + (((((0.0 + (float32(v1) * 0.0)) - 1e10) - v4) / v4) * (v4 * v4)))))))))
                _ = v39
            default:
                fmt.Println(v3, v29, v1)
                fmt.Println(fi37, v28)
                fmt.Println(v28)
            }
            fmt.Println(v28)
            _ = fi38
        }
        _ = fi37
    }
    if (4294967295 >= 2) {
        v28 = 127
        if (("\n\t" + (("a" + "") + (("hello" + ((("\n\t" + (("a" + ("a" + ("hello" + ""))) + (("\n\t" + "hello") + "a"))) + "a") + ("\n\t" + ""))) + "foo bar"))) == "hello") {
            v40 := []int{2147483647, 1}
            _ = v40
        } else {
            v41 := GenS3{F0: genFn1(), F1: 127}
            _ = v41
        }
        v42 := GenS5{F0: 194, F1: v28, F2: 4294967295}
        _ = v42
    }
    v43 := []int{(v3 & (v3 / (((^(v3))) | 1)))}
    var v44 uint8 = 1
    v1 = v1
    fmt.Println(v28, v28, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
