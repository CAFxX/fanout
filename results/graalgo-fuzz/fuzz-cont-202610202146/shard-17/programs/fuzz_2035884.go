package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint }

func main() {
    var v1 string = "a"
    var v2 uint16 = 2
    var v3 int8 = -1
    fmt.Println(v1, v3, v3)
    if false {
        var v4 int = 9223372036854775807
        _ = v4
    }
    v1 = "hello"
    v5 := []int{-760, -569}
    switch -591 {
    case 32768:
        v1 = ((v1 + ((v1 + "hello") + "a")) + v1)
    case 9223372036854775807:
        v3 = (v3 - int8(v2))
    case 128:
        var v6 float64 = 3.14159
        _ = v6
    }
    fmt.Println(v3, v1)
    if (-32768 >= 32767) {
        if true {
            for fi7 := 0; fi7 < 18; fi7++ {
                v1 = v1
                _ = fi7
            }
        } else {
            if true {
                if (true && true) {
                    var v8 bool = false
                    v2 -= v2
                    _ = v8
                }
            }
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int32 = 0
                _ = len(v5)
                _ = fi9
                _ = v10
            }
            v5 = append(v5, 1)
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            v2 = 2
            _ = fi11
        }
        if (true && ((!(!(true || (!((false && (!(!(((v2 & v2) < 255) && true)))) && true))))) && (965 == 42))) {
            for fi12 := 0; fi12 < 15; fi12++ {
                v1 = v1
                v1 = v1
                v3 = (v3 ^ (v3 ^ (v3 >> ((uint(v2)) & 63))))
                _ = fi12
            }
            v13 := GenS1{F0: -1, F1: 1, F2: 2}
            for fi14 := 0; fi14 < 19; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v2 = v2
                    v16 := GenS1{F0: 537, F1: 846, F2: 993}
                    _ = len(v5)
                    _ = fi15
                    _ = v16
                }
                switch fi14 {
                case 117:
                    if true {
                        var v17 bool = true
                        for fi18 := 0; fi18 < 11; fi18++ {
                            switch (fi18 << ((4294967295) & 63)) {
                            case 128:
                                if false {
                                    v5 = append(v5, (fi14 | (fi14 % ((int(v2)) | 1))))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v19 := GenS2{F0: 178, F1: 65535}
                                    _ = v19
                                }
                            default:
                                var v20 float64 = 1.5
                                _ = v20
                            }
                            v21 := GenS2{F0: 0, F1: 838}
                            var v22 int = (fi14 % ((-1) | 1))
                            _ = fi18
                            _ = v21
                            _ = v22
                        }
                        var v23 int8 = 42
                        _ = v17
                        _ = v23
                    }
                }
                _ = fi14
            }
            _ = v13
        } else {
            for fi24 := 0; fi24 < 15; fi24++ {
                for fi25 := 0; fi25 < 12; fi25++ {
                    v26 := GenS1{F0: -876, F1: -1, F2: 1}
                    if len(v5) > 0 {
                        v5[0] = fi25
                    }
                    _ = fi25
                    _ = v26
                }
                v1 = (("hello" + v1) + ("hello" + "hello"))
                _ = fi24
            }
            _ = len(v5)
            if (1.5 < 0.1) {
                v27 := GenS1{F0: 9223372036854775807, F1: 385, F2: 4294967295}
                _ = v27
            } else {
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 uint8 = 42
                    v30 := GenS2{F0: 1, F1: 42}
                    v31 := GenS3{F0: 42}
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                var v32 uint = 2
                switch v1 {
                case "hello":
                    fmt.Println(v32, v1, v2)
                case "\n\t":
                    v2 -= v2
                default:
                    v3 -= 1
                }
                _ = v32
            }
        }
    }
    var v33 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v5
}
