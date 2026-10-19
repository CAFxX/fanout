package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0() uint8 {
    var v6 int64 = -9223372036854775808
    if true {
        for fi7 := 0; fi7 < 7; fi7++ {
            v6 = 1
            _ = fi7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint16 = 32767
            v9 = v9
            _ = fi8
            _ = v9
        }
    } else {
        if false {
            for fi10 := 0; fi10 < 12; fi10++ {
                if (-1 == 2147483647) {
                    switch (fi10 >> ((255) & 63)) {
                    case -58:
                        v11 := []int{(fi10 % (((fi10 - 32767)) | 1))}
                        _ = v11
                    case 46:
                        v6 = 42
                    }
                    var v12 uint16 = 0
                    v12 += 42
                    _ = v12
                }
                v13 := []int{}
                _ = fi10
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 1}
                var v16 float64 = 1e10
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        v17 := []int{32768}
        var v18 int = -7
        _ = v17
        _ = v18
    }
    _ = v6
    return 255
}
func genFn1() int16 {
    var v19 uint8 = 1
    v20 := GenS1{F0: 0}
    v21 := GenS1{F0: 255}
    for fi22 := 0; fi22 < 9; fi22++ {
        switch fi22 {
        case 856:
            switch fi22 {
            case 32768:
                fmt.Println(fi22, fi22)
            default:
                fmt.Println(v19, v19)
            }
        case 32768:
            if (!false) {
                var v23 int32 = -1
                _ = v23
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    fmt.Println(fi24, fi24, fi24)
                    fmt.Println(v19, fi24, v19)
                    var v25 int16 = 1000
                    switch 128 {
                    case 9223372036854775807:
                        v26 := []int{32768, (fi24 & fi24), (fi22 >> ((0) & 63)), fi22}
                        v19 -= 1
                        var v27 uint = 18446744073709551615
                        v28 := GenS2{F0: 1, F1: "\n\t", F2: 32767}
                        _ = v26
                        _ = v27
                        _ = v28
                    case 42:
                        var v29 uint16 = 255
                        _ = v29
                    default:
                        var v30 uint = 1
                        _ = v30
                    }
                    v31 := GenS2{F0: 664, F1: "", F2: 386}
                    _ = fi24
                    _ = v25
                    _ = v31
                }
            }
        }
        fmt.Println(fi22, v19)
        _ = fi22
    }
    v19 = genFn0()
    _ = v19
    _ = v20
    _ = v21
    return 32767
}

func main() {
    var v1 int64 = 42
    var v2 string = ("" + "")
    var v3 int16 = 42
    var v4 float32 = 3.14159
    var v5 string = v2
    if true {
        v4 = v4
        v5 = v2
    } else {
        var v32 float32 = float32(float64(v4))
        _ = v32
    }
    fmt.Println(v2)
    if false {
        v3 += genFn1()
        v33 := GenS1{F0: 987}
        _ = v33
    } else {
        var v34 int = 42
        _ = v34
    }
    var v35 uint8 = 0
    var v36 uint = 1
    switch ("a" + (v5 + "")) {
    case "\n\t":
        for fi37 := 0; fi37 < 20; fi37++ {
            fmt.Println(fi37, v1)
            v38 := GenS1{F0: 493}
            _ = fi37
            _ = v38
        }
    case "":
        var v39 int = 32768
        _ = v39
    }
    fmt.Println(v3, v3, v3)
    switch "foo bar" {
    case "a":
        v2 = v5
    }
    var v40 uint16 = 32767
    v5 = ""
    fmt.Println(v36, v36, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
}
