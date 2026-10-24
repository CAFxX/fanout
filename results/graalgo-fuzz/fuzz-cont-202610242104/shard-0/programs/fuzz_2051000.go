package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 bool }
func genFn0() int {
    var v6 uint32 = 1
    v6 = v6
    _ = v6
    return 128
}

func main() {
    v1 := []string{"a"}
    var v2 uint16 = 32767
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    var v7 float64 = -2.25
    v8 := GenS1{F0: 4294967295, F1: 1000, F2: (v2 << ((18446744073709551615) & 63))}
    var v9 int8 = 0
    if true {
        var v10 float32 = 1.5
        _ = v10
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 bool = true
            fmt.Println(v7, v9)
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            v9 = v9
            _ = fi13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int = -516
            var v16 uint64 = 2
            v17 := GenS1{F0: uint32(v2), F1: 1000, F2: (v2 + 2)}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    v9 = v9
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            v7 += v7
            fmt.Println(v9)
            _ = fi19
        }
        _ = fi18
    }
    v9 = 127
    var v20 int16 = 1000
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            if ((v7 * v7) > 3.14159) {
                var v23 uint = 383
                delete(v3, 2147483647)
                var v24 float32 = (float32(float64(3.14159)) + 1e10)
                _ = v23
                _ = v24
            } else {
                v1 = append(v1, "hello")
            }
            switch (fi21 % ((fi21) | 1)) {
            case 42:
                switch ("" + ("a" + "")) {
                case "\n\t":
                    if true {
                        switch 226 {
                        case 521:
                            if false {
                                v25 := GenS1{F0: 65535, F1: (v20 * (v20 % (((v20 | v20)) | 1))), F2: 134}
                                v26 := v5["a"]
                                if (false && ((!false) || true)) {
                                    v27 := GenS1{F0: 65535, F1: (v20 - v20), F2: (v2 + (^(v2)))}
                                    var v28 bool = (!true)
                                    v7 = (v7 / v7)
                                    if v28 {
                                        if true {
                                            v2 += v2
                                        } else {
                                            _ = len(v4)
                                        }
                                    } else {
                                        switch fi22 {
                                        case -1:
                                            v29 := GenS2{F0: fi22, F1: 255, F2: 0}
                                            _ = v29
                                        }
                                    }
                                    v30 := GenS1{F0: 978, F1: 32767, F2: v2}
                                    var v31 string = v26
                                    _ = v27
                                    _ = v28
                                    _ = v30
                                    _ = v31
                                }
                                _ = v25
                                _ = v26
                            }
                        case -802:
                            v7 = float64(3.14159)
                        default:
                            v7 += v7
                        }
                    }
                case "":
                    switch (("foo bar" + "hello") + "hello") {
                    case "a":
                        if false {
                            delete(v3, 2147483648)
                        }
                    default:
                        var v32 uint32 = 65535
                        _ = v32
                    }
                case "hello":
                    var v33 string = "a"
                    v34 := GenS3{F0: 65535, F1: 18446744073709551615, F2: true}
                    _ = v33
                    _ = v34
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if true {
                    var v35 int16 = (v20 / ((105) | 1))
                    v4[("hello" + "hello")] = fi22
                    _ = v35
                }
            }
            v2 = 42
            _ = fi22
        }
        switch (fi21 + (^(fi21))) {
        case 42:
            var v36 uint = 2
            _ = v36
        case 1:
            var v37 int = fi21
            _ = v37
        }
        _ = fi21
    }
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
