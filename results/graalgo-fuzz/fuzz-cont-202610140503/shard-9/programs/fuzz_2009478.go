package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int8 = 100
    var v2 float32 = (3.14159 + 3.14159)
    v3 := []uint32{558}
    v4 := make(map[int]int)
    var v5 int8 = v1
    v6 := make(map[uint32]string)
    switch ("hello" + ("a" + ("\n\t" + "hello"))) {
    case "a":
        if (false && false) {
            var v7 string = ("\n\t" + ("a" + "hello"))
            _ = v7
        } else {
            if (255 <= 0) {
                switch 9223372036854775807 {
                case -1:
                    v1 -= v5
                }
                v3 = append(v3, 0)
                var v8 int8 = v1
                _ = v8
            }
            var v9 uint64 = 319
            _ = v9
        }
    case "":
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                v1 = (v5 >> ((65535) & 63))
                _ = len(v4)
                _ = fi11
            }
            var v12 int = fi10
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS1{F0: 1, F1: -991, F2: 0}
                v2 = v2
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v12
        }
    case "hello":
        switch 1 {
        case 128:
            switch 128 {
            case -1:
                v15 := GenS2{F0: 9223372036854775807, F1: 42}
                _ = v15
            case -804:
                v16 := GenS2{F0: 42, F1: 864}
                _ = v16
            }
        case -2147483648:
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS3{F0: 543, F1: (("hello" + "a") + ("hello" + "\n\t"))}
                if false {
                    v2 = v2
                    v5 += -43
                    var v19 int64 = 1
                    _ = v19
                } else {
                    v20 := GenS1{F0: 286, F1: -957, F2: -1}
                    v21 := GenS4{F0: 18446744073709551615}
                    _ = v20
                    _ = v21
                }
                var v22 uint16 = 2
                _ = fi17
                _ = v18
                _ = v22
            }
        case 42:
            var v23 uint64 = 437
            _ = v23
        }
    }
    fmt.Println(v2, v5)
    if false {
        for fi24 := 0; fi24 < 2; fi24++ {
            if (!(!false)) {
                fmt.Println(v5)
                v25 := GenS2{F0: 753, F1: 2147483648}
                _ = v25
            }
            _ = fi24
        }
    } else {
        if true {
            var v26 float32 = v2
            _ = v26
        }
        v27 := GenS1{F0: 18, F1: 1, F2: 42}
        _ = v27
    }
    v1 += v5
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v28 := GenS2{F0: -1, F1: 42}
        _ = v28
    case "foo bar":
        v1 = (v5 * v1)
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            v6[45] = "\n\t"
            v2 -= v2
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            _ = fi30
        }
        delete(v4, fi29)
        if len(v3) > 0 {
            v3[0] = 0
        }
        _ = fi29
    }
    v2 = v2
    v31 := GenS2{F0: 96, F1: 0}
    var v32 int8 = v1
    v33 := GenS2{F0: -1, F1: 0}
    if (false || (!true)) {
        if (("a" + "hello") < ("\n\t" + "a")) {
            v34 := GenS2{F0: 0, F1: 1}
            _ = v34
        } else {
            var v35 uint32 = 2147483647
            v1 -= (v1 << ((42) & 63))
            var v36 bool = (!true)
            _ = v35
            _ = v36
        }
        _ = len(v6)
    }
    fmt.Println(v2)
    var v37 string = "foo bar"
    v37 = v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
