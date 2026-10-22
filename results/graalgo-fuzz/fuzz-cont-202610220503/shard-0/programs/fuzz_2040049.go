package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint16; F1 uint16; F2 string }
func genFn0(p0 uint16, p1 int64, p2 int8) uint32 {
    var v7 uint64 = uint64(p0)
    if true {
        v7 += (v7 << ((4294967295) & 63))
        var v8 int = 32767
        fmt.Println(v7, v8)
        _ = v8
    }
    p2 = (p2 + 1)
    var v9 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 65535
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 int = 903
    var v4 bool = (v1 != (^(v1)))
    v5 := make(map[int]string)
    var v6 int16 = -1
    var v10 bool = false
    v6 = v6
    v11 := []int{}
    v1 = v1
    var v12 int = 0
    v4 = v4
    for fi13 := 0; fi13 < 1; fi13++ {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi13
    }
    var v14 uint8 = 127
    var v15 uint8 = uint8(v1)
    if true {
        for fi16 := 0; fi16 < 20; fi16++ {
            if (!((("\n\t" + "hello") <= "hello") || v4)) {
                for fi17 := 0; fi17 < 1; fi17++ {
                    if v4 {
                        v18 := GenS1{F0: -1, F1: 0}
                        v19 := GenS2{F0: -1, F1: 1, F2: 1}
                        _ = len(v11)
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
            } else {
                if false {
                    switch ((("" + "") + ((("\n\t" + "") + "a") + (("foo bar" + ("\n\t" + "")) + "a"))) + "") {
                    case "\n\t":
                        if v4 {
                            var v20 float32 = 1e10
                            switch (fi16 * v12) {
                            case -2147483648:
                                var v21 string = "\n\t"
                                _ = v21
                            case 9223372036854775807:
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    v23 := v5[v12]
                                    var v24 string = v23
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                            case 0:
                                var v25 uint = (v1 & 1)
                                _ = v25
                            default:
                                if v10 {
                                    var v26 bool = ((v6 + v6) > (v6 ^ int16(v1)))
                                    fmt.Println(v6, v4, v4)
                                    _ = v26
                                }
                            }
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 string = "a"
                                var v29 uint32 = 65535
                                fmt.Println(v14, v4, v29)
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            _ = v20
                        } else {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                v12 = fi16
                                _ = fi30
                            }
                            switch 128 {
                            case 128:
                                if (!false) {
                                    for fi31 := 0; fi31 < 3; fi31++ {
                                        v4 = v10
                                        var v32 bool = true
                                        _ = fi31
                                        _ = v32
                                    }
                                    v6 -= -555
                                } else {
                                    v33 := GenS2{F0: -1, F1: 875, F2: 4294967295}
                                    _ = v33
                                }
                            case 32768:
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    var v35 uint32 = 521
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            var v36 int16 = v6
                            _ = v36
                        }
                    case "":
                        fmt.Println(fi16, v12, v15)
                    case "hello":
                        var v37 uint32 = genFn0(774, 821, 100)
                        _ = v37
                    }
                    var v38 int = v12
                    _ = v38
                } else {
                    var v39 float64 = (0.1 + (0.0 * (((1e10 - 1e10) * -2.25) - (float64(1e10) - (((-2.25 - ((3.14159 / 0.1) + 3.14159)) / (1e10 / (-2.25 + 1.5))) + float64(0.0))))))
                    for fi40 := 0; fi40 < 4; fi40++ {
                        v1 += 0
                        if true {
                            if (v10 && (42 != -562)) {
                                v41 := v2[2147483648]
                                _ = v41
                            } else {
                                v12 += fi40
                                v14 += v15
                            }
                            if (v12 >= fi16) {
                                v42 := GenS3{F0: 0, F1: 1}
                                _ = v42
                            } else {
                                v6 -= (v6 - v6)
                                var v43 string = ("foo bar" + "\n\t")
                                _ = v43
                            }
                        } else {
                            v44 := GenS2{F0: 1, F1: 42, F2: 2}
                            v39 = 1e10
                            var v45 string = "hello"
                            _ = v44
                            _ = v45
                        }
                        var v46 uint = (^(v1))
                        if (!(false && v10)) {
                            if ("hello" <= ("a" + "foo bar")) {
                                if (("hello" + "") <= "foo bar") {
                                    var v47 uint32 = genFn0(2, -654, 1)
                                    var v48 float32 = 3.14159
                                    _ = v47
                                    _ = v48
                                }
                                var v49 float64 = v39
                                _ = v49
                            }
                            var v50 int64 = -448
                            var v51 uint8 = uint8(v1)
                            _ = v50
                            _ = v51
                        } else {
                            var v52 float32 = 0.0
                            var v53 int16 = v6
                            if true {
                                if (v52 > float32(0.1)) {
                                    var v54 uint64 = 1
                                    _ = v54
                                }
                                v55 := GenS4{F0: ((("\n\t" + "foo bar") + "") + "hello"), F1: v12}
                                _ = v55
                            }
                            _ = v52
                            _ = v53
                        }
                        _ = fi40
                        _ = v46
                    }
                    _ = v39
                }
                var v56 int8 = 127
                _ = v56
            }
            var v57 bool = (0 != 255)
            switch fi16 {
            case 127:
                var v58 float32 = (0.0 * (((-2.25 - 1e10) * 0.0) / (3.14159 * float32(3.14159))))
                _ = v58
            }
            _ = fi16
            _ = v57
        }
    }
    var v59 uint16 = 0
    var v60 float64 = (0.0 + 1.5)
    fmt.Println(v10, v4)
    for fi61 := 0; fi61 < 15; fi61++ {
        v59 -= v59
        v3 += (^(v12))
        v62 := GenS5{F0: (^(v59)), F1: v59, F2: "a"}
        _ = fi61
        _ = v62
    }
    fmt.Println(v59, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
