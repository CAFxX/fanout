package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int = 2147483647
    var v2 bool = true
    var v3 int8 = 42
    var v4 uint64 = 18446744073709551615
    var v5 uint8 = 2
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 1}
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 65535
    v4 = (v4 + v4)
    v9 := GenS1{F0: 42}
    v4 += 1
    var v10 int16 = 1
    if (-9223372036854775808 <= v1) {
        switch "foo bar" {
        case "":
            var v11 uint = 827
            _ = v11
        case "foo bar":
            switch "foo bar" {
            case "hello":
                v8 = (v8 % (((v8 % (((v8 & (v8 + 131))) | 1))) | 1))
            }
        default:
            switch (v1 * (v1 >> ((252) & 63))) {
            case 1:
                v8 = (v8 + v8)
            case 32768:
                if ((uint32(v4) <= v8) && (!(false && v2))) {
                    if (1.5 <= (float64(1.5) + 1.5)) {
                        v10 += v10
                        var v12 int8 = v3
                        _ = v12
                    }
                    if true {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 uint32 = v8
                            switch "" {
                            case "a":
                                v15 := []int{}
                                _ = v15
                            case "foo bar":
                                for fi16 := 0; fi16 < 7; fi16++ {
                                    var v17 uint8 = 255
                                    _ = fi16
                                    _ = v17
                                }
                            case "\n\t":
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    if false {
                                        if (378 < 388) {
                                            v19 := []int{}
                                            var v20 uint32 = (v8 + (^(v8)))
                                            var v21 uint8 = (v5 * 255)
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        } else {
                                            var v22 float32 = 0.0
                                            var v23 float32 = (float32((1.5 + 3.14159)) * v22)
                                            v14 = (v14 / ((0) | 1))
                                            _ = v22
                                            _ = v23
                                        }
                                        v14 -= 0
                                        v4 = 18446744073709551615
                                    } else {
                                        v24 := GenS1{F0: int16(v5)}
                                        v2 = (-747 != v1)
                                        _ = v24
                                    }
                                    v25 := []int{}
                                    _ = fi18
                                    _ = v25
                                }
                            default:
                                for fi26 := 0; fi26 < 13; fi26++ {
                                    var v27 int32 = int32(v4)
                                    v28 := GenS1{F0: v10}
                                    if v2 {
                                        v29 := GenS1{F0: v10}
                                        _ = v29
                                    }
                                    v30 := []int{v1, v1, 4294967295}
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v30
                                }
                            }
                            _ = fi13
                            _ = v14
                        }
                        v31 := GenS1{F0: v10}
                        _ = v31
                    }
                }
            default:
                for fi32 := 0; fi32 < 19; fi32++ {
                    if true {
                        v33 := []int{443}
                        _ = v33
                    }
                    v8 = v8
                    _ = fi32
                }
            }
        }
    } else {
        for fi34 := 0; fi34 < 12; fi34++ {
            if (0 >= 157) {
                v35 := GenS1{F0: 1}
                _ = v35
            }
            var v36 int32 = 42
            _ = fi34
            _ = v36
        }
    }
    switch (("hello" + "\n\t") + "") {
    case "foo bar":
        if (!(("\n\t" + ("a" + "")) > "hello")) {
            var v37 string = "a"
            _ = v37
        } else {
            v38 := []int{v1, 9223372036854775807, 32767}
            _ = v38
        }
    case "":
        switch (v1 & (v1 & v1)) {
        case -894:
            v10 = v10
        default:
            v2 = ((v2 || v2) && (false && (uint64(v5) != v4)))
        }
    default:
        if false {
            if (-2147483648 != -1) {
                var v39 int = (v1 + v1)
                _ = v39
            } else {
                fmt.Println(v4, v5, v1)
            }
        } else {
            fmt.Println(v5, v10)
        }
    }
    v40 := GenS1{F0: (v10 + (v10 / ((v10) | 1)))}
    v2 = (-586 >= 42)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v8
    _ = v9
}
