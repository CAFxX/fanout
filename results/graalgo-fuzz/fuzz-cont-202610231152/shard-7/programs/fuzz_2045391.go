package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 42
    var v3 string = ("a" + "hello")
    v4 := []uint32{42, 42, 223}
    v5 := []string{"a", v3, v3, "a"}
    for fi6 := 0; fi6 < 15; fi6++ {
        switch fi6 {
        case -498:
            var v7 string = v3
            _ = v7
        }
        var v8 uint16 = 890
        var v9 int16 = 32767
        _ = fi6
        _ = v8
        _ = v9
    }
    var v10 uint8 = 128
    fmt.Println(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint16 = 65535
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint8 = (v10 >> ((2) & 63))
            var v15 uint = 17
            if len(v4) > 0 {
                v4[0] = 964
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v10)
    fmt.Println(v10)
    v16 := GenS1{F0: 9223372036854775807, F1: v10}
    var v17 string = v3
    if (v10 < (v10 + 4)) {
        _ = len(v1)
        fmt.Println(v17)
    }
    v10 += (v10 + (v10 + (v10 >> ((65535) & 63))))
    v1[42] = (("\n\t" + v3) + (v3 + "foo bar"))
    fmt.Println(v3, v10, v10)
    _ = len(v4)
    v18 := GenS1{F0: 1, F1: 138}
    v3 = ("hello" + v3)
    switch v17 {
    case "\n\t":
        var v19 uint32 = 712
        _ = v19
    case "hello":
        for fi20 := 0; fi20 < 10; fi20++ {
            v3 = (("foo bar" + ("a" + "a")) + (v17 + (v3 + (v17 + v17))))
            _ = fi20
        }
    }
    if true {
        if true {
            switch (v17 + (("foo bar" + (v3 + v3)) + v3)) {
            case "foo bar":
                var v21 uint64 = 0
                _ = v21
            case "a":
                for fi22 := 0; fi22 < 10; fi22++ {
                    v23 := GenS1{F0: 170, F1: v10}
                    _ = fi22
                    _ = v23
                }
            }
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 bool = (-32768 < -32768)
            if len(v4) > 0 {
                v4[0] = 25
            }
            v2 += v2
            _ = fi24
            _ = v25
        }
        var v26 bool = false
        _ = v26
    } else {
        for fi27 := 0; fi27 < 7; fi27++ {
            if true {
                var v28 int64 = 1
                v10 += 163
                _ = v28
            }
            v29 := v1[fi27]
            _ = fi27
            _ = v29
        }
        if true {
            if (!true) {
                var v30 uint64 = 0
                _ = v30
            }
            var v31 uint32 = 1
            _ = v31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
