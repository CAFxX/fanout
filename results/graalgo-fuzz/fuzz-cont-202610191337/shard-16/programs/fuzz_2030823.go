package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 string; F1 int16 }
func genFn0() uint {
    var v7 uint32 = 1
    v8 := []int{128}
    v9 := GenS1{F0: -1}
    _ = v7
    _ = v8
    _ = v9
    return 469
}

func main() {
    var v1 uint16 = 8
    var v2 int32 = 42
    var v3 int16 = 818
    v4 := make(map[uint32]uint32)
    v5 := []string{"hello"}
    var v6 bool = (0 >= 358)
    var v10 bool = (v2 == 42)
    var v11 int64 = 819
    v12 := GenS1{F0: 1}
    if (!(true || v10)) {
        v1 -= (^(v1))
        var v13 int = 32768
        _ = v13
    } else {
        v14 := GenS1{F0: 1}
        if v6 {
            var v15 int = -976
            v1 = (^(v1))
            v4[0] = 1
            _ = v15
        }
        _ = v14
    }
    v16 := GenS2{F0: 911, F1: int(v3), F2: 678}
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint32 = 709
        if false {
            var v19 uint32 = (v18 % (((v18 | v18)) | 1))
            var v20 int64 = -150
            _ = v19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v21 := GenS2{F0: 42, F1: -1, F2: 2}
    switch -184 {
    case 32768:
        v3 = (v3 + 32767)
    case -935:
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS3{F0: "\n\t"}
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 bool = false
                _ = fi24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
    }
    if (!v6) {
        delete(v4, 42)
    } else {
        if (int16(v2) < (^(v3))) {
            if v6 {
                var v26 string = ""
                _ = v26
            }
            v27 := GenS4{F0: 624, F1: "\n\t"}
            _ = v27
        } else {
            if false {
                v28 := GenS5{F0: (("" + "a") + ""), F1: (v3 * v3)}
                switch "hello" {
                case "a":
                    for fi29 := 0; fi29 < 18; fi29++ {
                        if (v3 == v3) {
                            v3 += v3
                            var v30 float64 = 1e10
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v30
                        } else {
                            var v31 uint32 = 1
                            v10 = v6
                            _ = v31
                        }
                        fmt.Println(fi29, v10, v6)
                        var v32 int16 = (v3 | (v3 | v3))
                        _ = fi29
                        _ = v32
                    }
                case "hello":
                    var v33 string = ""
                    _ = v33
                }
                if ("hello" > "a") {
                    var v34 int16 = v3
                    delete(v4, 528)
                    _ = v34
                }
                _ = v28
            } else {
                for fi35 := 0; fi35 < 8; fi35++ {
                    v36 := v4[404]
                    var v37 int64 = 1
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                v11 = 9223372036854775807
            }
        }
        if true {
            var v38 uint = 42
            _ = v38
        }
        v4[4294967295] = 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
