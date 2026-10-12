package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 bool }
func genFn0(p0 int32, p1 int8, p2 uint) uint32 {
    v4 := []int{0}
    p1 += 57
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}
func genFn1(p0 int, p1 uint16) int32 {
    switch "a" {
    case "":
        for fi5 := 0; fi5 < 2; fi5++ {
            switch ("a" + (("\n\t" + "a") + ("a" + "hello"))) {
            case "":
                var v6 bool = false
                _ = v6
            case "hello":
                fmt.Println(p1, p0)
            case "a":
                switch "hello" {
                case "\n\t":
                    fmt.Println(p0, fi5)
                case "hello":
                    v7 := []int{}
                    _ = v7
                }
            }
            var v8 int = (p0 / ((998) | 1))
            _ = fi5
            _ = v8
        }
    }
    if ("\n\t" >= ("a" + "foo bar")) {
        if (0 < 105) {
            for fi9 := 0; fi9 < 10; fi9++ {
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(p0)
                    p1 = p1
                    _ = fi10
                }
                _ = fi9
            }
        } else {
            var v11 int64 = 692
            _ = v11
        }
        var v12 int = (p0 << ((42) & 63))
        for fi13 := 0; fi13 < 11; fi13++ {
            switch v12 {
            case -9223372036854775808:
                var v14 float32 = 0.0
                _ = v14
            default:
                v15 := GenS1{F0: 554, F1: 525}
                _ = v15
            }
            v16 := []int{}
            _ = fi13
            _ = v16
        }
        _ = v12
    }
    v17 := []int{p0, 127, (p0 + (^(p0)))}
    _ = p0
    _ = p1
    _ = v17
    return 1
}

func main() {
    var v1 uint16 = 2
    var v2 bool = true
    v3 := make(map[string]uint32)
    v18 := []int{0, -485, 4294967295}
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(v1, v2)
        _ = fi19
    }
    v20 := v3["a"]
    var v21 float64 = -2.25
    var v22 uint16 = v1
    _ = len(v18)
    fmt.Println(v22, v22)
    fmt.Println(v2, v20, v21)
    var v23 float64 = ((v21 / (float64(float32(float64(1.5))) + v21)) * 1e10)
    var v24 bool = (753 != 42)
    v25 := GenS1{F0: (v20 - genFn0(genFn1(608, (v1 - (v1 + v22))), 0, 2)), F1: 42}
    var v26 int = 753
    for fi27 := 0; fi27 < 5; fi27++ {
        v2 = true
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        if v24 {
            if v2 {
                for fi28 := 0; fi28 < 16; fi28++ {
                    switch (fi27 >> ((507) & 63)) {
                    case -1:
                        if v2 {
                            fmt.Println(v1, v26, v21)
                            v22 = v22
                            var v29 uint32 = (v20 + (v20 * (v20 % (((v20 * v20)) | 1))))
                            _ = v29
                        }
                    }
                    var v30 int64 = 1
                    var v31 int8 = -35
                    fmt.Println(v2, v23)
                    switch (fi28 % ((fi28) | 1)) {
                    case -218:
                        var v32 string = "a"
                        _ = v32
                    default:
                        var v33 uint64 = 503
                        _ = v33
                    }
                    var v34 int16 = -32768
                    if v2 {
                        v35 := GenS1{F0: v20, F1: 42}
                        var v36 uint = 2
                        _ = v35
                        _ = v36
                    } else {
                        switch fi28 {
                        case 32768:
                            fmt.Println(v22, fi27, v1)
                        case 0:
                            v31 += (v31 >> ((42) & 63))
                        case 485:
                            if (42 < 1) {
                                switch "a" {
                                case "hello":
                                    v37 := GenS2{F0: 42}
                                    _ = v37
                                case "":
                                    v18 = append(v18, (fi27 + fi27))
                                    delete(v3, "")
                                default:
                                    var v38 int64 = (v30 * v30)
                                    _ = v38
                                }
                                switch fi28 {
                                case 0:
                                    v1 += (^(v1))
                                case 32768:
                                    v21 -= v23
                                case -1:
                                    var v39 int16 = -32768
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    v40 := GenS2{F0: 839}
                                    _ = len(v18)
                                    fmt.Println(v26, v31, v24)
                                    var v41 bool = (!false)
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                            } else {
                                v18 = append(v18, 668)
                                v42 := GenS3{F0: (^(v1)), F1: (!v24)}
                                v43 := GenS3{F0: v22, F1: ((-2147483648 == -752) && true)}
                                _ = v42
                                _ = v43
                            }
                        default:
                            if v2 {
                                v44 := GenS4{F0: 4294967295}
                                v30 = v30
                                _ = v44
                            } else {
                                v30 = v30
                                v1 -= v1
                            }
                        }
                    }
                    fmt.Println(v24, fi27, fi27)
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    v23 = (((v23 + (v21 * v23)) + float64(0.0)) + (1.5 / (v23 * v23)))
                    _ = fi28
                    _ = v30
                    _ = v31
                    _ = v34
                }
                v45 := GenS3{F0: v22, F1: ((v20 - uint32(v1)) != (v20 & (v20 * v20)))}
                v23 -= v23
                _ = v45
            }
            for fi46 := 0; fi46 < 1; fi46++ {
                fmt.Println(v23, v22)
                v47 := GenS5{F0: 100, F1: (1 < -2147483648)}
                switch fi46 {
                case -315:
                    var v48 uint8 = 42
                    _ = v48
                case 42:
                    var v49 int8 = 42
                    _ = v49
                }
                _ = fi46
                _ = v47
            }
        }
        _ = fi27
    }
    v22 = 255
    var v50 float64 = v23
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v50
}
