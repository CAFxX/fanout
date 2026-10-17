package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    var v3 int = -814
    v4 := make(map[uint32]uint32)
    var v5 bool = false
    var v6 int64 = 0
    var v7 uint8 = 2
    var v8 uint32 = 2147483648
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: "foo bar"}
        v11 := GenS1{F0: "a"}
        if ((2147483647 > 42) && ((float64(1e10) / 1.0) > 0.1)) {
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS1{F0: (((("hello" + ((("a" + (("foo bar" + ("foo bar" + "")) + "")) + ((((("\n\t" + "a") + "") + ("" + (("hello" + "") + "foo bar"))) + "hello") + "a")) + "foo bar")) + "\n\t") + "a") + ("foo bar" + "hello"))}
                _ = fi12
                _ = v13
            }
            v4[v8] = (v8 % ((65535) | 1))
            if ("a" != "foo bar") {
                for fi14 := 0; fi14 < 8; fi14++ {
                    if (1000 <= 703) {
                        if (-9223372036854775808 > v6) {
                            v15 := []int{}
                            v7 = v7
                            v2 -= v2
                            _ = v15
                        } else {
                            v16 := []int{fi14}
                            _ = v16
                        }
                    } else {
                        v17 := GenS1{F0: "foo bar"}
                        v2 = float32(0.1)
                        if v5 {
                            delete(v1, (fi9 / (((fi9 >> ((2) & 63))) | 1)))
                        }
                        _ = v17
                    }
                    v18 := GenS1{F0: (("" + "hello") + "foo bar")}
                    v3 -= (fi14 << ((1) & 63))
                    _ = fi14
                    _ = v18
                }
            }
        } else {
            v8 = (^(v8))
            switch (fi9 % ((-797) | 1)) {
            case -35:
                if (v7 > (^(v7))) {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v6, v7)
                    case "hello":
                        fmt.Println(v6)
                    case "foo bar":
                        var v19 int32 = 988
                        _ = v19
                    }
                } else {
                    switch fi9 {
                    case 32768:
                        for fi20 := 0; fi20 < 9; fi20++ {
                            var v21 uint64 = 0
                            _ = fi20
                            _ = v21
                        }
                    }
                }
            case -927:
                for fi22 := 0; fi22 < 19; fi22++ {
                    var v23 uint64 = 0
                    if (false && (!v5)) {
                        switch fi9 {
                        case -2147483648:
                            v8 = v8
                            var v24 int64 = -1
                            if (v23 < (v23 & 2)) {
                                v7 += v7
                                var v25 uint8 = (v7 >> ((1) & 63))
                                _ = v25
                            }
                            v26 := []int{4294967295, 32767, 564}
                            v8 += v8
                            _ = v24
                            _ = v26
                        case 42:
                            v6 = (v6 >> ((1) & 63))
                        case 32768:
                            var v27 float64 = float64(v2)
                            _ = v27
                        default:
                            var v28 float32 = v2
                            _ = v28
                        }
                        fmt.Println(v8, fi22)
                    } else {
                        var v29 int8 = -128
                        _ = v29
                    }
                    v3 = (fi9 + (v3 << ((0) & 63)))
                    _ = fi22
                    _ = v23
                }
            case 32768:
                for fi30 := 0; fi30 < 6; fi30++ {
                    var v31 uint = 1
                    _ = fi30
                    _ = v31
                }
            }
            for fi32 := 0; fi32 < 19; fi32++ {
                v33 := v1[-465]
                v7 -= (^(v7))
                _ = fi32
                _ = v33
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v34 uint64 = 18446744073709551615
    v1[(-(v3))] = (("a" + ("\n\t" + "\n\t")) + ("hello" + ("a" + "foo bar")))
    v35 := GenS1{F0: (((("" + "") + "hello") + "\n\t") + "")}
    v6 = (v6 | 42)
    v3 -= (v3 - int(v7))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
