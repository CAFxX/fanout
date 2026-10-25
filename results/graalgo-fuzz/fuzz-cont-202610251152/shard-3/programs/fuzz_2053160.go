package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint16 }
func genFn0(p0 uint64) int8 {
    var v6 uint = 18446744073709551615
    v7 := []int{2147483647}
    v6 -= v6
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int8) uint64 {
    p0 -= p0
    var v8 uint64 = 18446744073709551615
    switch 861 {
    case 68:
        p0 = (p0 - (-(p0)))
    case 874:
        switch -485 {
        case 128:
            var v9 int8 = (p0 | p0)
            _ = v9
        case 292:
            switch "" {
            case "":
                var v10 string = ""
                _ = v10
            case "a":
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 int64 = 1
                    _ = fi11
                    _ = v12
                }
            default:
                var v13 int = 42
                _ = v13
            }
        case 0:
            for fi14 := 0; fi14 < 6; fi14++ {
                if (!(0.0 >= (1e10 / 0.1))) {
                    var v15 uint32 = 42
                    fmt.Println(v15, fi14, v8)
                    _ = v15
                }
                v8 += v8
                _ = fi14
            }
        }
    case -118:
        var v16 uint32 = 2
        _ = v16
    default:
        v17 := GenS1{F0: 1, F1: uint16(v8), F2: genFn0((^(v8)))}
        _ = v17
    }
    v18 := []int{32768, -2147483648}
    var v19 uint8 = 1
    _ = p0
    _ = v18
    _ = v19
    _ = v8
    return (v8 ^ 476)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-425, 2147483647, -2147483648, 0}
    v3 := []uint32{}
    var v4 int16 = -428
    v5 := []string{(("hello" + "") + "a"), "foo bar", ("\n\t" + ("" + "a")), "hello"}
    _ = len(v5)
    v4 = v4
    v20 := GenS1{F0: 65535, F1: 32767, F2: -1}
    var v21 float32 = ((0.0 - 1.5) - (float32(-2.25) - 1.5))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v22 := GenS1{F0: 849, F1: 2, F2: -1}
    switch 32768 {
    case -411:
        if true {
            if false {
                v23 := GenS2{F0: 787, F1: ""}
                var v24 uint8 = 255
                _ = v23
                _ = v24
            }
            switch 480 {
            case 2147483647:
                for fi25 := 0; fi25 < 16; fi25++ {
                    fmt.Println(v4)
                    fmt.Println(v4, v4, fi25)
                    fmt.Println(v4, v21)
                    _ = fi25
                }
            case -9223372036854775808:
                v4 -= -32768
            case -1:
                v4 -= v4
            }
            fmt.Println(v4, v4)
        } else {
            v26 := GenS2{F0: 298, F1: (("a" + ("" + "foo bar")) + "\n\t")}
            if true {
                v27 := GenS2{F0: 0, F1: "hello"}
                v4 = (v4 - -727)
                v28 := GenS1{F0: 18446744073709551615, F1: 42, F2: 56}
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 uint32 = 51
                    switch fi29 {
                    case 128:
                        v1["a"] = 2147483647
                    }
                    var v31 uint32 = v30
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                v32 := v1["foo bar"]
                var v33 float32 = float32((1.5 - 0.1))
                _ = v32
                _ = v33
            }
            _ = v26
        }
    case -1:
        var v34 uint = 65535
        _ = v34
    case -9223372036854775808:
        var v35 uint64 = genFn1(-109)
        _ = v35
    }
    if false {
        for fi36 := 0; fi36 < 3; fi36++ {
            switch "\n\t" {
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                fmt.Println(v21)
            }
            if (false && true) {
                fmt.Println(v21, v4, fi36)
                v37 := v1["foo bar"]
                _ = v37
            } else {
                for fi38 := 0; fi38 < 11; fi38++ {
                    fmt.Println(v21)
                    if true {
                        var v39 string = "hello"
                        switch v39 {
                        case "":
                            var v40 float64 = 0.0
                            _ = v40
                        case "hello":
                            var v41 string = ""
                            _ = v41
                        case "a":
                            v42 := GenS3{F0: 42, F1: 2, F2: 32767}
                            _ = v42
                        default:
                            fmt.Println(v39, fi38)
                        }
                        var v43 bool = (genFn0(1) < 1)
                        _ = v39
                        _ = v43
                    } else {
                        var v44 int64 = -255
                        v1[(("foo bar" + "hello") + (("a" + "hello") + "foo bar"))] = 0
                        fmt.Println(fi36)
                        _ = v44
                    }
                    _ = fi38
                }
                if false {
                    fmt.Println(v21, v21)
                    fmt.Println(v4, v4, v21)
                }
                v45 := GenS1{F0: 18446744073709551615, F1: 65535, F2: genFn0(742)}
                _ = v45
            }
            var v46 float32 = v21
            _ = fi36
            _ = v46
        }
        var v47 uint32 = 65535
        _ = v47
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v48 := GenS1{F0: 704, F1: 2, F2: 118}
        _ = v48
    }
    delete(v1, "foo bar")
    if false {
        var v49 int16 = v4
        if true {
            for fi50 := 0; fi50 < 2; fi50++ {
                v51 := GenS3{F0: genFn1(genFn0(4294967295)), F1: 2147483648, F2: 953}
                fmt.Println(v4, v21)
                _ = fi50
                _ = v51
            }
        }
        for fi52 := 0; fi52 < 2; fi52++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "foo bar" {
            case "a":
                v4 = -1
            case "\n\t":
                v21 += float32(((0.1 * 0.0) + (0.1 - float64((0.1 + (v21 * v21))))))
            case "hello":
                var v53 uint32 = 2147483647
                _ = v53
            default:
                v4 += (-(v4))
            }
            _ = fi52
        }
        _ = v49
    }
    if true {
        if true {
            v4 = 1
        } else {
            fmt.Println(v21, v21)
            for fi54 := 0; fi54 < 5; fi54++ {
                v55 := GenS1{F0: 255, F1: 2, F2: 107}
                fmt.Println(v4, v21, fi54)
                _ = fi54
                _ = v55
            }
            var v56 uint16 = 2
            _ = v56
        }
        var v57 uint32 = 42
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
