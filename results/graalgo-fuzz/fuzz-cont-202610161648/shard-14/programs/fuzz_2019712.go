package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32; F1 uint16; F2 int16 }
func genFn0(p0 uint64, p1 uint, p2 uint8) int16 {
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 bool = true
            fmt.Println(fi4)
            _ = fi5
            _ = v6
        }
        v7 := []int{(^(fi4)), -9223372036854775808, (fi4 * fi4)}
        _ = fi4
        _ = v7
    }
    if true {
        v8 := []int{127, 1, -586}
        fmt.Println(p2, p0)
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: fi9, F1: 127}
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: fi9, F1: 127}
            switch "\n\t" {
            case "hello":
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 string = "\n\t"
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = len(v8)
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                if (!((1 > p2) || true)) {
                    var v15 uint64 = p0
                    _ = v15
                }
            }
            _ = fi9
            _ = v12
        }
        _ = v8
    } else {
        var v16 uint64 = 814
        var v17 float32 = 3.14159
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return -912
}
func genFn1(p0 uint64, p1 int16, p2 uint64) int64 {
    var v18 uint64 = p2
    v19 := []int{-9223372036854775808, 540}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"hello"}
    v3 := make(map[int]string)
    var v20 bool = false
    switch (("a" + "a") + "") {
    case "\n\t":
        var v21 int32 = -1
        _ = v21
    case "a":
        if len(v2) > 0 {
            v2[0] = (("\n\t" + "hello") + "foo bar")
        }
    }
    if v20 {
        v20 = (!(false || v20))
    }
    if (!v20) {
        switch 4294967295 {
        case 2147483647:
            switch "a" {
            case "hello":
                v22 := GenS1{F0: 32768, F1: 1}
                _ = v22
            }
        case 801:
            if (42 < 87) {
                var v23 string = "a"
                _ = v23
            } else {
                var v24 bool = v20
                if v24 {
                    var v25 uint8 = 111
                    switch 4294967295 {
                    case 4294967295:
                        var v26 uint16 = 2
                        _ = v26
                    default:
                        for fi27 := 0; fi27 < 2; fi27++ {
                            switch (fi27 - 4294967295) {
                            case 9223372036854775807:
                                var v28 uint32 = 2147483647
                                _ = v28
                            case 1:
                                v29 := v3[fi27]
                                _ = v29
                            case 0:
                                if (!true) {
                                    var v30 uint32 = 805
                                    var v31 uint = 49
                                    fmt.Println(fi27)
                                    _ = v30
                                    _ = v31
                                }
                            }
                            for fi32 := 0; fi32 < 19; fi32++ {
                                if true {
                                    var v33 int32 = 2147483647
                                    _ = v33
                                } else {
                                    var v34 uint8 = 0
                                    _ = v34
                                }
                                v2 = append(v2, "hello")
                                _ = fi32
                            }
                            for fi35 := 0; fi35 < 11; fi35++ {
                                if true {
                                    v20 = v20
                                    var v36 float32 = 3.14159
                                    v25 += (v25 | v25)
                                    var v37 uint32 = 1
                                    v37 += v37
                                    _ = v36
                                    _ = v37
                                }
                                v24 = v24
                                var v38 string = (("hello" + "hello") + "hello")
                                fmt.Println(fi27)
                                v39 := v3[-842]
                                _ = fi35
                                _ = v38
                                _ = v39
                            }
                            _ = fi27
                        }
                    }
                    _ = v25
                }
                _ = v24
            }
        }
        if true {
            var v40 uint8 = 255
            switch "a" {
            case "\n\t":
                fmt.Println(v20, v40)
            default:
                var v41 uint16 = 1
                _ = v41
            }
            _ = v40
        } else {
            var v42 uint64 = 0
            _ = v42
        }
        var v43 uint32 = 2147483648
        _ = v43
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v44 int16 = 32767
    var v45 int32 = 0
    v20 = (255 == 255)
    v20 = ("a" < ("hello" + ("a" + "")))
    for fi46 := 0; fi46 < 7; fi46++ {
        v47 := v1[0]
        v45 = v45
        v48 := GenS1{F0: (fi46 + 32767), F1: 100}
        _ = fi46
        _ = v47
        _ = v48
    }
    for fi49 := 0; fi49 < 9; fi49++ {
        v50 := GenS2{F0: v45, F1: 94, F2: v44}
        v51 := GenS1{F0: (fi49 % ((9223372036854775807) | 1)), F1: 42}
        var v52 uint64 = 2
        _ = fi49
        _ = v50
        _ = v51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v44
    _ = v45
}
