package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 int8 }
type GenS4 struct { F0 string; F1 int }
func genFn0() int {
    if false {
        v5 := []int{127, 2147483647, 127}
        var v6 float64 = 0.0
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint = 1
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    switch 741 {
    case 2147483647:
        v9 := []int{}
        _ = v9
    case -1:
        var v10 string = "foo bar"
        _ = v10
    }
    if (!(-1 > -1)) {
        v11 := GenS1{F0: -641, F1: 431, F2: 9223372036854775807}
        _ = v11
    }
    v12 := GenS1{F0: -1, F1: 65535, F2: 710}
    switch "hello" {
    case "a":
        for fi13 := 0; fi13 < 13; fi13++ {
            if (1 <= 42) {
                if false {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        var v15 uint16 = 65535
                        v16 := []int{fi14, 2147483647, fi14, 411}
                        _ = len(v16)
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                } else {
                    var v17 uint = 65535
                    _ = v17
                }
                var v18 float32 = 1.5
                _ = v18
            }
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 string = "a"
                v21 := []int{374}
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 uint8 = 2
                _ = fi22
                _ = v23
            }
            _ = fi13
        }
    case "":
        v24 := GenS1{F0: 2147483647, F1: 922, F2: -50}
        _ = v24
    }
    _ = v12
    return 128
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1
    var v3 uint32 = 4294967295
    var v4 string = "foo bar"
    if (987 <= 1) {
        delete(v1, (v4 + v4))
    } else {
        v25 := GenS2{F0: v2, F1: uint64(v3)}
        _ = v25
    }
    if false {
        v4 = v4
        v3 -= (^(v3))
        if true {
            var v26 string = ("hello" + v4)
            _ = v26
        }
    }
    if true {
        var v27 string = v4
        v27 = v4
        _ = v27
    } else {
        if ((42 <= 42) || true) {
            if true {
                fmt.Println(v2, v2)
                switch ((((v4 + ("foo bar" + "hello")) + (v4 + ((((v4 + v4) + ((("" + "") + (v4 + ((("" + v4) + "hello") + "foo bar"))) + v4)) + ((v4 + "") + ((v4 + v4) + (v4 + "\n\t")))) + v4))) + ("hello" + (v4 + ("foo bar" + "\n\t")))) + "hello") {
                case "a":
                    for fi28 := 0; fi28 < 10; fi28++ {
                        switch v4 {
                        case "hello":
                            if true {
                                switch (fi28 * genFn0()) {
                                case 127:
                                    v1["a"] = (v4 + ((("hello" + (v4 + "foo bar")) + (v4 + "foo bar")) + ((v4 + ("" + "")) + (v4 + ("foo bar" + (v4 + v4))))))
                                case 128:
                                    v4 = ((("a" + ("a" + v4)) + v4) + (v4 + (v4 + ("\n\t" + ""))))
                                case -2147483648:
                                    v3 = (v3 | (v3 | 2147483647))
                                default:
                                    var v29 string = v4
                                    _ = v29
                                }
                                for fi30 := 0; fi30 < 7; fi30++ {
                                    var v31 bool = (("hello" + v4) == "a")
                                    var v32 uint64 = 18446744073709551615
                                    v33 := GenS3{F0: 100, F1: 943, F2: 113}
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                v34 := []int{fi28, (-(fi28)), 970, fi28}
                                _ = v34
                            }
                        case "":
                            v35 := GenS4{F0: v4, F1: genFn0()}
                            _ = v35
                        case "foo bar":
                            var v36 uint8 = 90
                            _ = v36
                        default:
                            if ((^(v3)) < (^(v3))) {
                                for fi37 := 0; fi37 < 8; fi37++ {
                                    fmt.Println(v3)
                                    var v38 float32 = -2.25
                                    fmt.Println(v4, v38)
                                    _ = fi37
                                    _ = v38
                                }
                                v3 -= 0
                                v3 += 812
                            } else {
                                var v39 uint8 = 1
                                _ = v39
                            }
                        }
                        _ = fi28
                    }
                }
                v2 += (v2 << ((65535) & 63))
            }
        }
    }
    v40 := []int{445, -1, 9223372036854775807, -9223372036854775808}
    fmt.Println(v4, v2)
    var v41 float32 = ((0.0 / 1e10) + float32((0.1 * (3.14159 + (0.0 * 1.5)))))
    var v42 uint16 = 2
    v4 = "a"
    v2 = (^(v2))
    var v43 int32 = -1
    var v44 int16 = 50
    if ((v3 | (v3 << ((746) & 63))) >= v3) {
        var v45 int16 = 1000
        v4 = (("hello" + v4) + v4)
        v1["hello"] = "foo bar"
        _ = v45
    } else {
        var v46 uint16 = 2
        _ = v46
    }
    v43 = (^(v43))
    v4 = "hello"
    if len(v40) > 0 {
        v40[0] = 32768
    }
    if true {
        v4 = (v4 + ((v4 + v4) + v4))
    }
    for fi47 := 0; fi47 < 14; fi47++ {
        v4 = v4
        _ = fi47
    }
    v1[(v4 + ("hello" + ""))] = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
