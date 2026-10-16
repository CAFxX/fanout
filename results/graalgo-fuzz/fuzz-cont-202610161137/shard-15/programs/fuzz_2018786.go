package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool }
func genFn0() uint64 {
    if true {
        v4 := []int{}
        var v5 int16 = 891
        fmt.Println(v5)
        _ = v4
        _ = v5
    } else {
        if (-9223372036854775808 > 127) {
            switch "" {
            case "hello":
                if false {
                    for fi6 := 0; fi6 < 16; fi6++ {
                        switch fi6 {
                        case 32767:
                            var v7 uint16 = 2
                            _ = v7
                        case 856:
                            var v8 bool = (-1 == 941)
                            _ = v8
                        case 42:
                            var v9 string = ("a" + "foo bar")
                            _ = v9
                        }
                        _ = fi6
                    }
                    for fi10 := 0; fi10 < 1; fi10++ {
                        fmt.Println(fi10)
                        if true {
                            var v11 uint16 = 1
                            _ = v11
                        }
                        _ = fi10
                    }
                } else {
                    var v12 uint16 = 65535
                    if false {
                        var v13 int = -891
                        for fi14 := 0; fi14 < 8; fi14++ {
                            fmt.Println(v12)
                            var v15 uint16 = v12
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    } else {
                        switch 1 {
                        case -2147483648:
                            for fi16 := 0; fi16 < 14; fi16++ {
                                var v17 int16 = 338
                                v18 := []int{fi16, fi16, (fi16 ^ fi16), (fi16 ^ fi16)}
                                v12 -= v12
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        default:
                            if (4294967295 > 65535) {
                                for fi19 := 0; fi19 < 5; fi19++ {
                                    v12 = 323
                                    fmt.Println(fi19, fi19)
                                    _ = fi19
                                }
                                for fi20 := 0; fi20 < 6; fi20++ {
                                    switch ("" + (("" + "\n\t") + "")) {
                                    case "\n\t":
                                        v21 := []int{int(v12), fi20, (^(fi20))}
                                        _ = v21
                                    }
                                    _ = fi20
                                }
                                v12 += v12
                            } else {
                                if ("a" == "") {
                                    fmt.Println(v12)
                                    v12 -= (v12 / (((v12 << ((4294967295) & 63))) | 1))
                                } else {
                                    var v22 uint8 = 1
                                    var v23 int16 = -1
                                    _ = v22
                                    _ = v23
                                }
                            }
                        }
                    }
                    switch ("" + "foo bar") {
                    case "\n\t":
                        var v24 int = -1
                        _ = v24
                    case "":
                        v25 := GenS1{F0: 42}
                        _ = v25
                    default:
                        v26 := GenS1{F0: 42}
                        _ = v26
                    }
                    _ = v12
                }
            default:
                var v27 int16 = 42
                _ = v27
            }
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 uint64 = 42
                var v30 uint8 = 0
                v31 := []int{fi28, fi28}
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
        }
    }
    var v32 bool = true
    _ = v32
    return 18446744073709551615
}

func main() {
    var v1 int64 = -327
    v2 := []uint32{795, 2147483648}
    v3 := make(map[string]uint32)
    for fi33 := 0; fi33 < 10; fi33++ {
        if false {
            if true {
                v1 += (v1 * (v1 ^ (v1 | (v1 / ((9223372036854775807) | 1)))))
                var v34 uint8 = 3
                _ = v34
            } else {
                v1 -= (^(v1))
            }
            v3["\n\t"] = 42
        } else {
            var v35 string = ""
            _ = v35
        }
        var v36 float32 = ((1.5 - 0.0) * (1.5 / 1.0))
        switch (("\n\t" + "foo bar") + "\n\t") {
        case "foo bar":
            if false {
                _ = len(v3)
                var v37 int = fi33
                _ = v37
            }
        case "\n\t":
            var v38 uint64 = genFn0()
            _ = v38
        }
        _ = fi33
        _ = v36
    }
    switch (("hello" + "a") + "\n\t") {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 1
        }
    case "\n\t":
        v2 = append(v2, 305)
    case "hello":
        delete(v3, "a")
    default:
        v1 = v1
    }
    v2 = append(v2, 65535)
    var v39 uint16 = 42
    if false {
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 uint32 = 2147483648
            _ = fi40
            _ = v41
        }
    }
    switch 128 {
    case 4294967295:
        v39 -= v39
    }
    switch 2147483647 {
    case 653:
        v42 := GenS2{F0: true, F1: false}
        _ = v42
    }
    v1 += 69
    for fi43 := 0; fi43 < 16; fi43++ {
        v1 += (v1 * (v1 % (((v1 * -9223372036854775808)) | 1)))
        _ = fi43
    }
    var v44 bool = true
    fmt.Println(v39, v1)
    for fi45 := 0; fi45 < 19; fi45++ {
        var v46 int32 = int32(v1)
        _ = fi45
        _ = v46
    }
    var v47 uint32 = 265
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v44
    _ = v47
}
