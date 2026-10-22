package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 int32; F1 int32 }
func genFn0(p0 uint, p1 int8) uint16 {
    v6 := GenS1{F0: 198, F1: 32767, F2: (^(p0))}
    p0 += (p0 >> ((561) & 63))
    _ = p0
    _ = p1
    _ = v6
    return 2
}
func genFn1() int32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{(^(fi7)), 0}
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 42
    if true {
        v9 = v9
        v10 := GenS2{F0: 108}
        v11 := []int{-1, -2147483648}
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            v9 += 932
            _ = fi12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            if (!(false && false)) {
                v14 := GenS2{F0: 42}
                v15 := []int{(^(fi13))}
                _ = v14
                _ = v15
            } else {
                switch 0 {
                case 0:
                    v16 := GenS3{F0: 368, F1: 147, F2: 2147483647}
                    _ = v16
                case -68:
                    for fi17 := 0; fi17 < 4; fi17++ {
                        fmt.Println(fi17, fi17)
                        _ = fi17
                    }
                case -2147483648:
                    v18 := GenS4{F0: -1}
                    _ = v18
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 int64 = 42
                    var v21 int8 = -80
                    v20 = (v20 % (((v20 % (((v20 ^ 613)) | 1))) | 1))
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                fmt.Println(v9, fi13)
            }
            if (false || (2 == 13)) {
                fmt.Println(fi13, v9)
                if ((!((1e10 * (1e10 - (0.1 - ((0.0 / 1.0) * float64(3.14159))))) != 0.0)) || ((((fi13 >> ((931) & 63)) <= fi13) || false) && (!false))) {
                    v9 += (v9 - 4294967295)
                    if (!true) {
                        var v22 int32 = 852
                        v23 := []int{(fi13 % ((-2147483648) | 1))}
                        var v24 uint8 = 2
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS5{F0: (("" + "hello") == "")}
                    _ = v25
                } else {
                    fmt.Println(fi13, v9)
                    var v26 uint32 = (v9 << ((4294967295) & 63))
                    _ = v26
                }
            }
            _ = fi13
        }
        var v27 int64 = 0
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 uint16 = 42
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 string = ("" + ("a" + "foo bar"))
            var v32 int8 = 1
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    _ = v9
    return -308
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 int16 = -32768
    v4 := []int{1}
    v5 := make(map[int]string)
    for fi33 := 0; fi33 < 9; fi33++ {
        v34 := GenS3{F0: 65535, F1: 705, F2: fi33}
        fmt.Println(v3)
        _ = fi33
        _ = v34
    }
    var v35 uint = 690
    var v36 int = -807
    if true {
        fmt.Println(v3, v36, v36)
        var v37 uint8 = 198
        _ = v37
    } else {
        var v38 float64 = 0.0
        switch ((("a" + ("a" + "hello")) + "hello") + "") {
        case "a":
            v38 += (-2.25 + v38)
        case "foo bar":
            fmt.Println(v38, v3, v38)
        }
        var v39 float64 = v38
        _ = v38
        _ = v39
    }
    if len(v4) > 0 {
        v4[0] = (v36 - (-(v36)))
    }
    switch (v36 & 4294967295) {
    case -707:
        v35 += v35
    case 32767:
        switch v36 {
        case 349:
            v40 := GenS6{F0: 65535, F1: (v3 / ((v3) | 1))}
            _ = v40
        case -1:
            v41 := GenS5{F0: true}
            _ = v41
        case 0:
            var v42 uint64 = 311
            _ = v42
        }
    case 128:
        fmt.Println(v3)
    }
    v43 := GenS7{F0: 708, F1: 42}
    if (684 > 0) {
        v3 = (^(v3))
        v36 -= v36
        v44 := v2[459]
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v5
}
