package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int32) int32 {
    switch 529 {
    case 127:
        p0 = 2
    default:
        v4 := []int{32767, 2147483647, 0, 32768}
        _ = v4
    }
    v5 := GenS1{F0: 65535}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int32 = -2147483648
        p1 = (p1 + -1)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return (p1 << ((255) & 63))
}
func genFn1(p0 uint64, p1 int32, p2 int32) uint32 {
    v8 := GenS1{F0: 37}
    p1 += -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 string = ("\n\t" + "foo bar")
    v2 := make(map[string]uint32)
    var v3 uint16 = 65535
    delete(v2, v1)
    fmt.Println(v1)
    var v9 int64 = -1
    fmt.Println(v1, v1, v3)
    var v10 uint32 = 2147483647
    if (49 > -1) {
        if ("\n\t" == v1) {
            for fi11 := 0; fi11 < 13; fi11++ {
                switch (v1 + v1) {
                case "a":
                    var v12 string = ""
                    _ = v12
                }
                _ = fi11
            }
            var v13 float32 = ((0.1 - float32(1e10)) / -2.25)
            _ = v13
        }
        v9 = v9
    }
    v3 -= v3
    var v14 uint8 = 151
    v9 -= (^(v9))
    v9 = -9223372036854775808
    if true {
        switch -9223372036854775808 {
        case 0:
            if (true && true) {
                switch ("hello" + "hello") {
                case "a":
                    var v15 float64 = float64(0.0)
                    _ = v15
                case "foo bar":
                    var v16 uint = 0
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 string = (("\n\t" + "\n\t") + v1)
                        if true {
                            if true {
                                v19 := GenS1{F0: genFn1(4294967295, genFn0((v3 >> ((42) & 63)), 42), 1)}
                                switch (("\n\t" + "hello") + (v18 + "")) {
                                case "foo bar":
                                    fmt.Println(v10)
                                case "hello":
                                    v14 -= (v14 >> ((255) & 63))
                                case "":
                                    fmt.Println(fi17, v1, v1)
                                }
                                _ = v19
                            } else {
                                var v20 uint16 = (v3 * (v3 << ((65535) & 63)))
                                _ = v20
                            }
                            v18 = "hello"
                            var v21 uint8 = 128
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            }
        case -411:
            v3 += (^(v3))
        default:
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := []int{fi22, (fi22 - fi22)}
                var v24 uint8 = uint8(v3)
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        delete(v2, ((v1 + v1) + v1))
    }
    var v25 string = v1
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v9
}
