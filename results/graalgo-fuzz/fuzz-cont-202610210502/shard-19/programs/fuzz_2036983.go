package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
func genFn0(p0 int32, p1 uint32, p2 string) uint {
    v6 := []int{32767, int(p1), -921}
    v7 := GenS1{F0: false, F1: p1, F2: p2}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 115
}

func main() {
    v1 := []string{"a"}
    var v2 int8 = 0
    var v3 uint8 = 42
    v4 := make(map[uint32]bool)
    v5 := []uint32{2147483648, 2, 2147483647}
    var v8 int8 = -80
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint16 = 2
        _ = fi9
        _ = v10
    }
    v8 += v2
    var v11 uint = 255
    v11 = (^(v11))
    delete(v4, 361)
    v3 += (v3 << (((v11 * v11)) & 63))
    if false {
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                if (!(!true)) {
                    var v14 string = "foo bar"
                    _ = v14
                } else {
                    var v15 int32 = -433
                    _ = v15
                }
                if false {
                    fmt.Println(v11, v11, v2)
                    v3 -= v3
                    var v16 uint = genFn0(42, uint32(v11), "hello")
                    _ = v16
                } else {
                    if (127 > v3) {
                        v2 = v2
                    }
                    if true {
                        v3 = (v3 << ((v11) & 63))
                        if true {
                            var v17 int8 = v8
                            _ = v17
                        } else {
                            if len(v1) > 0 {
                                v1[0] = ("a" + "hello")
                            }
                            v3 += 2
                        }
                    }
                }
                _ = fi13
            }
            _ = fi12
        }
        var v18 int16 = -822
        _ = v18
    } else {
        v3 = 7
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v19 := GenS1{F0: (!(true || true)), F1: 65535, F2: ("\n\t" + (("" + "") + "a"))}
        _ = v19
    }
    var v20 float32 = float32((-2.25 / ((-2.25 + -2.25) + float64(3.14159))))
    v8 = 127
    if (true || (0.1 <= 3.14159)) {
        fmt.Println(v2, v11)
    } else {
        switch int(v8) {
        case -382:
            v21 := GenS1{F0: false, F1: 570, F2: "hello"}
            _ = v21
        case 1:
            v8 -= (v2 >> ((538) & 63))
        }
    }
    v3 -= 3
    switch ("foo bar" + "a") {
    case "\n\t":
        v5 = append(v5, 2)
    case "foo bar":
        var v22 int = 1
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
