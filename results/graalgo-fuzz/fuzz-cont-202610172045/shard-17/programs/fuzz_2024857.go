package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0(p0 uint8) int8 {
    var v6 int8 = 127
    v7 := GenS1{F0: 1}
    fmt.Println(v6, p0)
    switch ((("a" + (((((("\n\t" + "") + "a") + "\n\t") + "\n\t") + ("hello" + "hello")) + "hello")) + ("" + "foo bar")) + "\n\t") {
    case "hello":
        v8 := GenS1{F0: 509}
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return (v6 | v6)
}

func main() {
    var v1 uint32 = 2
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    var v4 int8 = 0
    var v5 uint = 1
    var v9 uint = v5
    for fi10 := 0; fi10 < 9; fi10++ {
        if v2 {
            v11 := GenS2{F0: "foo bar"}
            var v12 int8 = (v4 & v4)
            _ = v11
            _ = v12
        } else {
            var v13 int = (^(fi10))
            delete(v3, (v1 - (v1 ^ v1)))
            _ = v13
        }
        delete(v3, v1)
        _ = fi10
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v2 = ((-639 >= fi14) || false)
            if true {
                v16 := []int{(-(fi14)), fi14}
                _ = v16
            }
            _ = fi15
        }
        if true {
            var v17 uint16 = 2
            var v18 float64 = 1.5
            v19 := GenS1{F0: (v1 << ((42) & 63))}
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi14
    }
    switch 127 {
    case 4294967295:
        v20 := GenS1{F0: (v1 ^ v1)}
        _ = v20
    case 0:
        var v21 int = -9223372036854775808
        _ = v21
    case -2147483648:
        for fi22 := 0; fi22 < 18; fi22++ {
            if ((v1 >> ((v5) & 63)) < 65535) {
                if ((1e10 * float32(0.0)) >= 1.5) {
                    var v23 uint16 = 0
                    _ = v23
                }
                v9 = 874
            } else {
                fmt.Println(fi22)
            }
            v24 := GenS1{F0: (^(v1))}
            _ = fi22
            _ = v24
        }
    }
    switch (("\n\t" + (("\n\t" + ("\n\t" + "foo bar")) + "")) + "hello") {
    case "foo bar":
        switch -2147483648 {
        case 32768:
            v25 := GenS3{F0: -1}
            _ = v25
        case 974:
            fmt.Println(v2, v2, v2)
        case -220:
            var v26 string = (("a" + "\n\t") + "\n\t")
            _ = v26
        }
    case "a":
        switch ("a" + "") {
        case "foo bar":
            var v27 uint16 = 2
            _ = v27
        }
    }
    var v28 int = 9223372036854775807
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(v4)
            var v31 uint16 = 436
            _ = fi30
            _ = v31
        }
        v32 := GenS1{F0: 0}
        delete(v3, (v1 | (v1 | (^(v1)))))
        _ = fi29
        _ = v32
    }
    v33 := []int{9223372036854775807, v28}
    delete(v3, v1)
    switch ((("hello" + "foo bar") + "hello") + ("foo bar" + "hello")) {
    case "foo bar":
        var v34 bool = v2
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 uint8 = 0
            var v38 int16 = 0
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
