package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0(p0 uint32, p1 int32, p2 uint16) uint16 {
    var v4 uint16 = p2
    v5 := GenS1{F0: -1, F1: 12}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -299
    var v3 int = (v2 ^ (v2 ^ (v2 >> ((42) & 63))))
    var v6 int8 = -41
    var v7 uint32 = 490
    switch ((("a" + ("foo bar" + ("\n\t" + "\n\t"))) + "a") + (("a" + "a") + "\n\t")) {
    case "foo bar":
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint32 = v7
            v2 += v3
            _ = fi8
            _ = v9
        }
    case "":
        var v10 int32 = 443
        _ = v10
    }
    fmt.Println(v6, v2, v2)
    v3 = int(v7)
    switch ("a" + "") {
    case "hello":
        if (2 > 42) {
            delete(v1, (v7 * v7))
            v6 += (v6 * v6)
            if ((!true) && true) {
                fmt.Println(v2, v2)
                if true {
                    if true {
                        for fi11 := 0; fi11 < 6; fi11++ {
                            v3 += fi11
                            v1[(^(v7))] = "\n\t"
                            _ = fi11
                        }
                    } else {
                        v1[v7] = "a"
                        v12 := GenS1{F0: -1, F1: (v6 % (((v6 | 127)) | 1))}
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 13; fi13++ {
                        if true {
                            switch ("\n\t" + ("\n\t" + ("hello" + ("hello" + ("hello" + "\n\t"))))) {
                            case "a":
                                switch "\n\t" {
                                case "hello":
                                    var v14 int = (v3 - (fi13 << ((65535) & 63)))
                                    _ = v14
                                case "a":
                                    v1[(v7 & (v7 << ((4294967295) & 63)))] = "foo bar"
                                case "":
                                    v15 := []int{(v3 | v3), fi13}
                                    _ = v15
                                }
                            case "\n\t":
                                var v16 int16 = 1
                                _ = v16
                            case "":
                                var v17 float32 = ((1.5 - (0.1 * (-2.25 * 3.14159))) / 1e10)
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 9; fi18++ {
                                var v19 float32 = 0.0
                                var v20 float64 = 1e10
                                var v21 float64 = v20
                                v22 := GenS2{F0: false, F1: (v2 | fi18)}
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                        fmt.Println(v6, fi13)
                        switch (-(fi13)) {
                        case 4294967295:
                            v23 := GenS3{F0: (fi13 * 2147483647), F1: 255, F2: 226}
                            _ = v23
                        default:
                            var v24 uint32 = v7
                            _ = v24
                        }
                        _ = fi13
                    }
                }
                v6 = v6
            } else {
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 uint8 = 127
                    _ = fi25
                    _ = v26
                }
                var v27 uint64 = 18446744073709551615
                v28 := GenS1{F0: 42, F1: -128}
                _ = v27
                _ = v28
            }
        } else {
            v29 := GenS3{F0: -1, F1: 255, F2: 1}
            _ = v29
        }
    }
    fmt.Println(v6, v6)
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := []int{int(v6), fi30, 0, -9223372036854775808}
        v32 := GenS2{F0: true, F1: (fi30 + (^(v2)))}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 float32 = (float32(3.14159) * float32(0.0))
    var v34 uint64 = 174
    var v35 uint8 = 255
    for fi36 := 0; fi36 < 1; fi36++ {
        v37 := GenS1{F0: 42, F1: v6}
        fmt.Println(v33, v6, v6)
        if true {
            v38 := GenS3{F0: (v2 * (v2 ^ v2)), F1: (v35 & 127), F2: (v35 % (((v35 >> ((57) & 63))) | 1))}
            _ = v38
        }
        _ = fi36
        _ = v37
    }
    var v39 int8 = (v6 + v6)
    var v40 uint8 = (v35 | (v35 % (((v35 + 1)) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v40
    _ = v6
    _ = v7
}
