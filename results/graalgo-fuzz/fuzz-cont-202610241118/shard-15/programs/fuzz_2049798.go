package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v1 += (v1 ^ 128)
    switch 1 {
    case 516:
        v1 += 0
    case -2147483648:
        v1 = (^(v1))
    case 1:
        if true {
            v4 := GenS1{F0: -13}
            if false {
                v1 = v1
            }
            v5 := GenS1{F0: -86}
            _ = v4
            _ = v5
        }
    default:
        v1 = (^(v1))
    }
    fmt.Println(v1)
    if (!false) {
        v1 = v1
        var v6 uint64 = 18446744073709551615
        v6 = v6
        _ = v6
    } else {
        v3[2147483647] = "\n\t"
    }
    if (true && (1 <= 32768)) {
        switch int(v1) {
        case 686:
            v7 := GenS1{F0: 0}
            _ = v7
        case 9223372036854775807:
            v8 := v3[uint32(v1)]
            _ = v8
        case 0:
            v1 -= (v1 | v1)
        default:
            v9 := GenS1{F0: 99}
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v1, v1)
            _ = fi10
        }
        var v11 bool = false
        v12 := []int{127, 4294967295}
        _ = v11
        _ = v12
    }
    switch "hello" {
    case "a":
        v13 := GenS1{F0: 50}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 19; fi14++ {
            v1 += v1
            v15 := []int{fi14}
            var v16 int32 = 2147483647
            _ = fi14
            _ = v15
            _ = v16
        }
    case "":
        for fi17 := 0; fi17 < 11; fi17++ {
            switch fi17 {
            case 32767:
                switch (fi17 >> ((360) & 63)) {
                case 789:
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v19 := v3[459]
                        _ = fi18
                        _ = v19
                    }
                case 42:
                    v1 = (v1 + 102)
                }
            case 2147483647:
                v20 := GenS1{F0: 100}
                _ = v20
            case 127:
                var v21 uint64 = 4294967295
                _ = v21
            default:
                if (!false) {
                    if false {
                        v1 = v1
                        fmt.Println(fi17, fi17)
                        var v22 uint = 0
                        _ = v22
                    }
                    var v23 int8 = -1
                    _ = v23
                }
            }
            fmt.Println(v1, v1)
            _ = fi17
        }
    }
    var v24 string = "hello"
    v24 = v24
    var v25 float32 = 0.0
    var v26 uint8 = (v1 - 128)
    switch "\n\t" {
    case "a":
        for fi27 := 0; fi27 < 11; fi27++ {
            fmt.Println(v25)
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 bool = true
                if v29 {
                    if (v29 && v29) {
                        v30 := []int{fi27}
                        _ = v30
                    }
                    var v31 float64 = 0.0
                    switch fi28 {
                    case 42:
                        var v32 uint32 = 295
                        _ = v32
                    case 32768:
                        var v33 uint = 1
                        var v34 uint16 = 42
                        var v35 uint8 = (v26 << (((v33 - v33)) & 63))
                        _ = v33
                        _ = v34
                        _ = v35
                    case -1:
                        v24 = ((v24 + ("" + v24)) + "foo bar")
                    default:
                        if false {
                            var v36 uint8 = (^(v1))
                            if true {
                                var v37 int64 = -1
                                var v38 int64 = v37
                                v39 := GenS1{F0: 42}
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            _ = v36
                        } else {
                            fmt.Println(v24)
                            var v40 float32 = v25
                            v41 := []int{fi27, 618, (fi28 << ((404) & 63))}
                            _ = v40
                            _ = v41
                        }
                    }
                    _ = v31
                }
                v42 := GenS1{F0: 100}
                _ = fi28
                _ = v29
                _ = v42
            }
            fmt.Println(v26, v26)
            _ = fi27
        }
    case "hello":
        _ = len(v2)
    }
    var v43 int = -690
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v43
}
