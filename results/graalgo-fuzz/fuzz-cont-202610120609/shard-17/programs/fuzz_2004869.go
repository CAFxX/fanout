package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 int64 = 872
    var v6 uint16 = 32767
    var v7 int8 = 42
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int32 = -2147483648
        _ = fi8
        _ = v9
    }
    switch (("" + ("hello" + "")) + "foo bar") {
    case "hello":
        v10 := GenS1{F0: 2147483647, F1: v5, F2: 42}
        _ = v10
    case "a":
        v5 -= v5
    default:
        if false {
            if false {
                var v11 uint8 = 2
                _ = v11
            } else {
                v7 += (v7 / ((v7) | 1))
                v5 = v5
                var v12 string = "\n\t"
                _ = v12
            }
            var v13 int32 = -308
            _ = v13
        } else {
            if (929 > (v6 + 65535)) {
                v6 += v6
                if false {
                    switch -1 {
                    case 32768:
                        if false {
                            var v14 float64 = 0.0
                            v15 := GenS1{F0: 71, F1: (v5 & (v5 & v5)), F2: 2}
                            _ = v14
                            _ = v15
                        }
                    case 9223372036854775807:
                        var v16 uint8 = 127
                        _ = v16
                    }
                }
                for fi17 := 0; fi17 < 14; fi17++ {
                    v6 = 0
                    v18 := GenS1{F0: -1, F1: 86, F2: 127}
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v20 := GenS1{F0: -1, F1: v5, F2: 110}
                        v6 = (^(v6))
                        _ = fi19
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
            }
            if (4294967295 >= 306) {
                var v21 int16 = -32768
                var v22 int32 = -1
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 14; fi23++ {
                    switch fi23 {
                    case -1:
                        var v24 int16 = int16(v5)
                        _ = v24
                    case 32768:
                        for fi25 := 0; fi25 < 9; fi25++ {
                            switch "hello" {
                            case "hello":
                                if ("foo bar" == ("\n\t" + "\n\t")) {
                                    delete(v2, 1)
                                    if false {
                                        var v26 int64 = (v5 | -1)
                                        v5 -= v5
                                        _ = v26
                                    }
                                    var v27 int64 = (v5 & v5)
                                    delete(v1, 855)
                                    _ = v27
                                } else {
                                    var v28 bool = true
                                    _ = v28
                                }
                            }
                            _ = fi25
                        }
                    }
                    v29 := GenS1{F0: 2147483647, F1: 42, F2: 42}
                    _ = fi23
                    _ = v29
                }
            }
            if (true && (127 != int(v6))) {
                for fi30 := 0; fi30 < 16; fi30++ {
                    if (42 > 756) {
                        v31 := []int{(^(fi30)), 4294967295, fi30}
                        if true {
                            var v32 uint32 = 913
                            _ = v32
                        }
                        _ = v31
                    }
                    var v33 bool = (498 != 42)
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v35 := v4[fi30]
                        fmt.Println(v5, v35)
                        var v36 float32 = 3.14159
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    _ = fi30
                    _ = v33
                }
            }
        }
    }
    var v37 bool = false
    v38 := GenS2{F0: v6, F1: 42}
    v39 := GenS3{F0: 485, F1: 127}
    v40 := []int{128, 2147483647, -9223372036854775808}
    delete(v2, 880)
    v5 += (v5 / ((9223372036854775807) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
