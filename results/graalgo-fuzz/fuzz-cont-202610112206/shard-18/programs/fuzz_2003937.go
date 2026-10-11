package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint = 1
    var v2 uint64 = 0
    v3 := make(map[int]string)
    var v4 float32 = 0.0
    var v5 int8 = -112
    var v6 uint8 = 196
    if (true || true) {
        fmt.Println(v6, v1, v2)
        var v7 uint16 = 2
        v1 += (v1 | v1)
        _ = v7
    } else {
        v8 := GenS1{F0: (("\n\t" + "foo bar") + "\n\t")}
        var v9 uint16 = 32767
        _ = v8
        _ = v9
    }
    fmt.Println(v2, v2)
    fmt.Println(v5, v4, v1)
    fmt.Println(v2, v1, v5)
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint16 = 2
        _ = fi10
        _ = v11
    }
    switch 4294967295 {
    case 32767:
        var v12 int64 = 1
        _ = v12
    default:
        for fi13 := 0; fi13 < 10; fi13++ {
            switch (((("a" + "\n\t") + ("hello" + "foo bar")) + ("hello" + "a")) + "foo bar") {
            case "foo bar":
                _ = len(v3)
            case "\n\t":
                var v14 uint32 = 2
                _ = v14
            case "hello":
                switch ("foo bar" + ("a" + "")) {
                case "hello":
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v16 := []int{(fi13 * fi13), fi13}
                        _ = fi15
                        _ = v16
                    }
                case "":
                    v4 -= (1.5 * v4)
                default:
                    if (uint16(v6) != 2) {
                        v17 := GenS1{F0: "a"}
                        fmt.Println(v4)
                        v2 = (v2 ^ v2)
                        _ = v17
                    } else {
                        v5 = -1
                    }
                }
            }
            _ = fi13
        }
    }
    v2 += v2
    for fi18 := 0; fi18 < 10; fi18++ {
        if false {
            if true {
                var v19 int16 = -1
                _ = v19
            }
            var v20 float32 = v4
            _ = v20
        } else {
            if true {
                var v21 int = -592
                for fi22 := 0; fi22 < 15; fi22++ {
                    if false {
                        var v23 int64 = 0
                        var v24 int = 32768
                        if (v1 >= (v1 * 65535)) {
                            var v25 string = ("\n\t" + (("a" + "") + "\n\t"))
                            fmt.Println(v23)
                            _ = v25
                        }
                        _ = v23
                        _ = v24
                    } else {
                        fmt.Println(v6, fi18)
                    }
                    var v26 uint8 = uint8(v1)
                    var v27 int64 = 1
                    _ = fi22
                    _ = v26
                    _ = v27
                }
                var v28 bool = (true && false)
                _ = v21
                _ = v28
            }
        }
        switch "\n\t" {
        case "\n\t":
            v4 = float32(float64(0.0))
        }
        _ = fi18
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := []int{fi29, fi29, 32767, -190}
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: (("hello" + "\n\t") + "")}
    fmt.Println(v5, v6, v5)
    if ("" > (("foo bar" + "\n\t") + (("a" + "") + "a"))) {
        var v32 uint = 2
        _ = v32
    } else {
        v1 = v1
        var v33 int16 = -1
        v1 -= v1
        _ = v33
    }
    var v34 int32 = -1
    if (14 >= 2) {
        v35 := []int{}
        v35 = append(v35, -650)
        _ = v35
    } else {
        var v36 uint16 = 2
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
