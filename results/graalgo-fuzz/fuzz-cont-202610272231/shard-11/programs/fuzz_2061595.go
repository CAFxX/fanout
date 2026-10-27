package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
func genFn0(p0 uint64) uint8 {
    var v7 int16 = 1000
    if false {
        v8 := []int{}
        if true {
            var v9 uint = 26
            _ = len(v8)
            switch 26 {
            case 128:
                var v10 float64 = (((3.14159 / 1.5) - -2.25) / 1.5)
                _ = v10
            case -2147483648:
                v11 := GenS1{F0: 2, F1: 0, F2: 1}
                _ = v11
            case 4294967295:
                fmt.Println(v9)
            }
            _ = v9
        }
        _ = v8
    } else {
        var v12 int8 = 0
        v13 := []int{127, 42, 0}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 80, F1: 65535, F2: 127}
    _ = p0
    _ = v14
    _ = v7
    return 255
}
func genFn1() uint32 {
    var v15 uint32 = 0
    fmt.Println(v15)
    v15 = (v15 << ((255) & 63))
    v16 := GenS1{F0: 42, F1: 931, F2: 605}
    v15 = 2
    _ = v15
    _ = v16
    return (^(v15))
}

func main() {
    var v1 uint = 406
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 uint16 = 645
    v17 := GenS1{F0: 255, F1: 2147483648, F2: 882}
    if ((!false) || ((0 < 42) || ((("\n\t" <= "\n\t") || (413 >= 67)) || (44 == -8)))) {
        v6 -= 798
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 bool = false
                _ = fi19
                _ = v20
            }
            if false {
                if false {
                    if false {
                        var v21 int64 = 9223372036854775807
                        if true {
                            var v22 int8 = -1
                            v21 = v21
                            _ = v22
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v23 uint32 = 2147483648
                            _ = v23
                        case "hello":
                            v24 := []int{fi18, fi18, fi18, 749}
                            _ = v24
                        }
                        _ = v21
                    } else {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            v26 := []int{(fi18 * (fi18 * -118)), (-(fi18)), fi18, fi18}
                            _ = fi25
                            _ = v26
                        }
                        var v27 uint8 = 1
                        _ = v27
                    }
                    v28 := GenS1{F0: 1, F1: 4294967295, F2: 128}
                    v29 := GenS1{F0: 189, F1: 2, F2: fi18}
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v1)
                }
            }
            _ = fi18
        }
    } else {
        var v30 int8 = 127
        v31 := GenS1{F0: 88, F1: genFn1(), F2: 0}
        var v32 bool = (0.1 >= 1e10)
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 int16 = -32768
    switch "foo bar" {
    case "a":
        var v34 string = ((("" + ((("a" + "a") + ("\n\t" + "\n\t")) + "hello")) + "hello") + "hello")
        _ = v34
    default:
        fmt.Println(v1)
    }
    v35 := v2[-2147483648]
    v6 -= 65535
    _ = len(v2)
    for fi36 := 0; fi36 < 2; fi36++ {
        v1 = (v1 - 1)
        v1 = v1
        _ = fi36
    }
    v37 := []int{9223372036854775807}
    for fi38 := 0; fi38 < 16; fi38++ {
        v1 -= v1
        _ = fi38
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        v6 = v6
        var v40 float32 = ((-2.25 - 3.14159) + (1.5 + 1e10))
        switch "hello" {
        case "":
            var v41 uint16 = (v6 & v6)
            _ = v41
        case "foo bar":
            var v42 uint = (v1 | (v1 + v1))
            _ = v42
        case "a":
            switch fi39 {
            case -1:
                if (true || (-83 >= 112)) {
                    if (395 <= -232) {
                        fmt.Println(v1)
                    }
                    v43 := v5[("foo bar" + "hello")]
                    _ = v43
                }
            }
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
