package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 string, p1 bool) uint32 {
    switch -697 {
    case 32768:
        p0 = (p0 + "foo bar")
    case 0:
        switch (p0 + (p0 + ((("a" + (p0 + p0)) + (p0 + p0)) + (p0 + "foo bar")))) {
        case "a":
            p0 = ""
        case "\n\t":
            v4 := GenS1{F0: 0, F1: 65535, F2: 100}
            _ = v4
        case "foo bar":
            fmt.Println(p0, p1)
        }
    case -482:
        v5 := []int{-9223372036854775808, 32768, -397}
        _ = v5
    default:
        switch "" {
        case "a":
            if (-9223372036854775808 >= 932) {
                switch 349 {
                case 1:
                    var v6 uint32 = 675
                    _ = v6
                default:
                    p1 = (-916 != 0)
                }
            } else {
                p1 = ((p0 + p0) >= p0)
                v7 := []int{1, 1}
                _ = v7
            }
        case "":
            if p1 {
                v8 := GenS2{F0: 2147483647, F1: 0}
                _ = v8
            }
        case "foo bar":
            p0 = "hello"
        }
    }
    v9 := []int{276}
    _ = len(v9)
    v10 := GenS2{F0: 920, F1: 32768}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 688
    v2 := make(map[string]string)
    var v3 float32 = ((1e10 + float32(0.1)) + (0.0 - float32(float64(0.1))))
    v3 += float32(1e10)
    var v11 int32 = 1
    v11 -= -1
    fmt.Println(v11)
    v1 = 4294967295
    if ((!(!(!true))) && (36 != 127)) {
        if (!((127 >= 0) && (1 == -57))) {
            var v12 uint = v1
            v11 = v11
            var v13 uint64 = 717
            _ = v12
            _ = v13
        } else {
            v14 := []int{}
            switch -384 {
            case 9223372036854775807:
                v3 = (v3 * v3)
            case -2147483648:
                var v15 bool = true
                _ = v15
            default:
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 string = ("hello" + "hello")
                    if (-851 < -9223372036854775808) {
                        switch v17 {
                        case "hello":
                            if false {
                                v3 -= (v3 / v3)
                                for fi18 := 0; fi18 < 8; fi18++ {
                                    var v19 bool = true
                                    _ = fi18
                                    _ = v19
                                }
                                v20 := GenS2{F0: 2, F1: fi16}
                                _ = v20
                            } else {
                                if false {
                                    var v21 bool = true
                                    _ = v21
                                }
                                for fi22 := 0; fi22 < 17; fi22++ {
                                    v2[v17] = v17
                                    if true {
                                        var v23 int = fi22
                                        _ = v23
                                    }
                                    fmt.Println(v11, v1, fi22)
                                    v24 := GenS1{F0: 54, F1: 2, F2: -75}
                                    _ = fi22
                                    _ = v24
                                }
                            }
                        default:
                            var v25 int32 = 0
                            _ = v25
                        }
                    }
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v27 := v2[(v17 + (("" + v17) + v17))]
                        var v28 uint64 = 2
                        switch 1 {
                        case 40:
                            if false {
                                switch fi16 {
                                case -774:
                                    var v29 int64 = int64(v11)
                                    _ = v29
                                case 9223372036854775807:
                                    v27 = "a"
                                }
                                if true {
                                    var v30 bool = false
                                    fmt.Println(v3)
                                    _ = v30
                                }
                                var v31 float64 = (1.5 - 1e10)
                                _ = v31
                            }
                        case 32768:
                            fmt.Println(fi26, fi26, v28)
                        default:
                            v32 := GenS3{F0: 1, F1: (!true)}
                            _ = v32
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = v14
        }
        var v33 uint32 = genFn0("\n\t", (1 >= 281))
        _ = v33
    }
    if ("a" < ("" + "a")) {
        v34 := GenS2{F0: genFn0("foo bar", (266 != v1)), F1: -700}
        var v35 uint8 = 2
        _ = v34
        _ = v35
    }
    v2[("hello" + "")] = "\n\t"
    for fi36 := 0; fi36 < 8; fi36++ {
        v37 := []int{fi36}
        delete(v2, "a")
        var v38 int32 = (-(v11))
        _ = fi36
        _ = v37
        _ = v38
    }
    if false {
        v11 -= 742
    }
    var v39 float64 = 0.0
    var v40 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}
