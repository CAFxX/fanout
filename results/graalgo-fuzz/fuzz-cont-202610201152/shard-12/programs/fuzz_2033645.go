package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0() int16 {
    v5 := []int{42, 9223372036854775807}
    _ = len(v5)
    _ = v5
    return 853
}
func genFn1(p0 bool) uint {
    p0 = (-32768 <= genFn0())
    p0 = true
    v6 := []int{-572, 42, -9223372036854775808}
    v7 := GenS1{F0: -1, F1: 659}
    var v8 uint = 65535
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 65535
    var v3 int8 = -112
    v4 := []int{-9223372036854775808, 4294967295, -9223372036854775808}
    switch 127 {
    case 597:
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint64 = 363
            _ = fi9
            _ = v10
        }
    case 2147483647:
        if false {
            var v11 int32 = -2147483648
            var v12 int = 128
            _ = len(v4)
            _ = v11
            _ = v12
        }
    case 1:
        delete(v1, "a")
    default:
        v2 = v2
    }
    v13 := GenS1{F0: -476, F1: 0}
    var v14 int64 = int64(v2)
    var v15 uint8 = 128
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            var v17 int8 = v3
            _ = v17
        } else {
            var v18 int8 = 100
            _ = v18
        }
        var v19 uint16 = (v2 * v2)
        _ = fi16
        _ = v19
    }
    fmt.Println(v2, v2, v14)
    switch "hello" {
    case "a":
        v20 := v1[("\n\t" + ("\n\t" + (("foo bar" + "hello") + "hello")))]
        _ = v20
    case "hello":
        v1[("hello" + "foo bar")] = 32767
    case "\n\t":
        if true {
            fmt.Println(v15, v3)
            if false {
                var v21 int16 = 1
                for fi22 := 0; fi22 < 7; fi22++ {
                    switch 2147483647 {
                    case 266:
                        for fi23 := 0; fi23 < 1; fi23++ {
                            if false {
                                switch fi22 {
                                case -167:
                                    delete(v1, (("hello" + "\n\t") + "foo bar"))
                                default:
                                    if (!false) {
                                        if true {
                                            v24 := GenS2{F0: 42}
                                            _ = v24
                                        }
                                        if false {
                                            v25 := GenS2{F0: 2}
                                            v26 := GenS3{F0: 4294967295, F1: 618, F2: v2}
                                            v2 = v2
                                            var v27 int = (fi22 + 128)
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        } else {
                                            fmt.Println(v21, v3, v15)
                                            v28 := GenS3{F0: 2, F1: 0, F2: v2}
                                            _ = v28
                                        }
                                    }
                                }
                                v1[""] = fi22
                                switch ("a" + "") {
                                case "a":
                                    switch int(v3) {
                                    case -171:
                                        v3 = -1
                                    }
                                case "":
                                    if (false && ((fi22 - fi23) >= fi22)) {
                                        var v29 int8 = 1
                                        _ = v29
                                    } else {
                                        v15 = (v15 + (v15 * 2))
                                        v15 -= v15
                                        switch fi23 {
                                        case -297:
                                            var v30 int8 = v3
                                            _ = v30
                                        case 4294967295:
                                            var v31 bool = (!(!(v2 != (v2 & (v2 * 255)))))
                                            _ = v31
                                        case -309:
                                            _ = len(v4)
                                        default:
                                            if true {
                                                fmt.Println(v14, v3, v3)
                                                fmt.Println(v14, v14)
                                                fmt.Println(v15, v2)
                                                v14 -= (v14 / (((^(v14))) | 1))
                                                _ = len(v4)
                                                v1["a"] = -474
                                            } else {
                                                v21 = v21
                                                v1[""] = fi23
                                                v21 = -32768
                                            }
                                        }
                                    }
                                }
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = (fi22 ^ fi23)
                                }
                                var v32 int32 = -2147483648
                                var v33 uint8 = v15
                                _ = v32
                                _ = v33
                            }
                            v34 := v1["\n\t"]
                            if len(v4) > 0 {
                                v4[0] = (fi23 % ((fi23) | 1))
                            }
                            _ = fi23
                            _ = v34
                        }
                    }
                    var v35 float32 = 1.5
                    _ = fi22
                    _ = v35
                }
                _ = v21
            }
            v2 += (v2 + v2)
        }
    default:
        if false {
            var v36 string = "foo bar"
            v2 = v2
            _ = v36
        }
    }
    var v37 uint8 = 255
    v3 += v3
    var v38 int32 = -2147483648
    var v39 int16 = genFn0()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
