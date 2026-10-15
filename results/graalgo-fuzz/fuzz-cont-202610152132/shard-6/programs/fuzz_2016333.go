package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int16; F2 uint8 }
func genFn0(p0 uint16) int64 {
    p0 = (^(p0))
    var v7 int32 = 42
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            switch fi8 {
            case 127:
                if false {
                    if (true || (!(-1 < v7))) {
                        v7 = (v7 / ((v7) | 1))
                    } else {
                        p0 += 272
                        if true {
                            for fi9 := 0; fi9 < 15; fi9++ {
                                var v10 bool = (float64(float32((3.14159 + (-2.25 * float64((3.14159 * (3.14159 + (float32((0.1 / 1.0)) * 0.0)))))))) != 1e10)
                                v11 := GenS1{F0: 126, F1: "hello"}
                                switch (-(fi8)) {
                                case -1:
                                    v12 := []int{(fi8 - fi9)}
                                    _ = v12
                                }
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            fmt.Println(v7, fi8, fi8)
                        }
                    }
                }
            case 128:
                p0 = (p0 >> ((2) & 63))
            case 4294967295:
                for fi13 := 0; fi13 < 15; fi13++ {
                    v14 := []int{(-(fi8)), fi13, int(p0), 0}
                    v7 = v7
                    v7 = (v7 % ((v7) | 1))
                    _ = fi13
                    _ = v14
                }
            }
            var v15 uint16 = (p0 << ((784) & 63))
            if false {
                for fi16 := 0; fi16 < 2; fi16++ {
                    if (127 <= 1) {
                        v17 := []int{(-(fi16)), (fi16 % ((fi16) | 1))}
                        v18 := GenS1{F0: 0, F1: "\n\t"}
                        v19 := GenS2{F0: "foo bar"}
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint = 1
                        fmt.Println(v15)
                        switch (("a" + "a") + "a") {
                        case "hello":
                            var v21 string = "\n\t"
                            _ = v21
                        case "foo bar":
                            v22 := []int{fi16, (fi16 - 128)}
                            _ = v22
                        case "a":
                            if false {
                                p0 = p0
                                var v23 uint16 = (^(v15))
                                v7 -= (-(v7))
                                var v24 string = "foo bar"
                                switch fi16 {
                                case 32767:
                                    if true {
                                        v7 = (v7 % (((v7 - (-(v7)))) | 1))
                                        v20 -= v20
                                        fmt.Println(fi8, v23)
                                    }
                                    v24 = "foo bar"
                                    v25 := GenS2{F0: v24}
                                    _ = v25
                                case -1:
                                    v23 = (^(p0))
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                        _ = v20
                    }
                    v7 -= v7
                    v15 = 32767
                    _ = fi16
                }
            }
            _ = fi8
            _ = v15
        }
        var v26 float32 = 1.5
        _ = v26
    } else {
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 string = "a"
            _ = fi27
            _ = v28
        }
        var v29 int = -314
        var v30 string = "foo bar"
        _ = v29
        _ = v30
    }
    var v31 string = "a"
    _ = p0
    _ = v31
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 bool = true
    v6 := make(map[int]string)
    v5 = v5
    var v32 int16 = -630
    var v33 uint64 = 1
    if (genFn0(32767) == 9223372036854775807) {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 uint64 = v33
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 uint8 = 255
        v32 += v32
        var v38 uint32 = 168
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint = 0
    var v40 int64 = 888
    switch ("" + "") {
    case "":
        v32 -= (v32 - (v32 + v32))
    default:
        v41 := GenS3{F0: 127, F1: v32, F2: 42}
        _ = v41
    }
    var v42 float64 = 0.1
    delete(v6, 1)
    var v43 int16 = v32
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
