package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 47
    v10 := GenS1{F0: 42}
    var v11 uint8 = 42
    _ = v10
    _ = v11
    _ = v9
    return -562
}
func genFn1(p0 uint8) string {
    v12 := GenS1{F0: -32768}
    v13 := []int{52, 127, genFn0()}
    _ = p0
    _ = v12
    _ = v13
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -13
    v3 := make(map[uint32]int)
    v4 := []uint32{1, 815, 673, 16}
    v5 := []uint32{1, 65535, 4294967295, 42}
    var v6 int = 657
    delete(v1, ((((("\n\t" + "a") + "") + ("foo bar" + ("hello" + ("a" + ((("hello" + "") + (("foo bar" + (("\n\t" + "") + ("hello" + ("foo bar" + ("\n\t" + ""))))) + "")) + ("\n\t" + "")))))) + ("" + "hello")) + ("" + "\n\t")))
    var v14 int64 = -471
    var v15 bool = true
    var v16 int32 = -1
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 bool = (1 == -773)
        _ = fi17
        _ = v18
    }
    fmt.Println(v14, v6, v6)
    switch v6 {
    case -822:
        var v19 uint32 = 42
        _ = v19
    case 922:
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 int16 = 32767
            _ = fi20
            _ = v21
        }
    case -1:
        switch (((("a" + "") + ((("a" + "foo bar") + "foo bar") + "a")) + "\n\t") + (("" + "\n\t") + "")) {
        case "\n\t":
            if v15 {
                switch v6 {
                case 4294967295:
                    var v22 uint8 = 255
                    _ = v22
                default:
                    var v23 int8 = 42
                    _ = v23
                }
                delete(v1, "")
            } else {
                var v24 int16 = -32768
                _ = v24
            }
        default:
            var v25 uint16 = 148
            _ = v25
        }
    }
    if (0 >= 2) {
        v14 -= v14
        var v26 uint32 = 1
        if false {
            v4 = append(v4, (v26 * v26))
            if ((-(v16)) < v16) {
                v16 += v16
                if (100 == (v2 / ((0) | 1))) {
                    switch "\n\t" {
                    case "":
                        delete(v1, ("foo bar" + "a"))
                    case "hello":
                        var v27 float64 = 0.0
                        _ = v27
                    case "a":
                        var v28 int64 = int64(v6)
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v2 = (v2 << ((65535) & 63))
                        v31 := GenS1{F0: 42}
                        fmt.Println(v16, v26, fi29)
                        _ = fi30
                        _ = v31
                    }
                    var v32 int32 = (v16 % (((v16 / ((v16) | 1))) | 1))
                    _ = fi29
                    _ = v32
                }
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v33 := v1["\n\t"]
                v34 := GenS1{F0: 32767}
                _ = v33
                _ = v34
            }
        }
        _ = v26
    }
    var v35 uint64 = 101
    var v36 uint16 = 140
    v15 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
